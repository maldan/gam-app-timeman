<template>
  <div class="history">
    <div class="block hour_map">
      <!-- Header -->
      <div class="header">Hour Map</div>

      <div class="time" v-for="(list, i) in hourMap" :key="i">
        <div class="hour">{{ ('0' + i).slice(-2) }}</div>
        <div class="bar">
          <div
            v-for="x in list"
            :key="x.offset"
            class="fill clickable"
            :style="{
              background: x.color,
              left: x.offset * 100 + '%',
              width: x.duration * 100 + '%',
            }"
          ></div>
          <div
            class="current"
            v-if="isToday && new Date().getHours() === i"
            :style="{ left: (new Date().getMinutes() / 60) * 100 + '%' }"
          >
            |
          </div>
        </div>
      </div>
    </div>

    <div class="block task_list">
      <!-- Header -->
      <div class="header">
        History
        <img @click="isShowAddForm = true" class="clickable" src="../../asset/add.svg" alt="" />
      </div>

      <div class="task" v-for="(x, i) in list" :key="x">
        <Task
          @edit="(isShowEditForm = true), (itemId = $event)"
          @delete="remove(x)"
          :item="x"
          :nextItem="list[i + 1]"
          :date="date"
        />
      </div>
    </div>

    <Add :date="date" v-if="isShowAddForm" @close="(isShowAddForm = false), $root.refresh()" />
    <Edit :id="itemId" v-if="isShowEditForm" @close="(isShowEditForm = false), $root.refresh()" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';
import Add from './Add.vue';
import Edit from './Edit.vue';
import Task from './Task.vue';
import Moment from 'moment';

export default defineComponent({
  props: {
    date: Object,
  },
  components: { Add, Edit, Task },
  async mounted() {
    this.color = (await RestApi.config.get())?.activityColor || {};
    this.refresh();
  },
  watch: {
    date(value: Date) {
      this.refresh();
    },
    '$root.refreshId'() {
      this.refresh();
    },
  },
  methods: {
    async refresh() {
      this.list = await RestApi.task.getByDay(Moment(this.date).format('YYYY-MM-DD'));
      this.hourMap.length = 0;
      for (let i = 0; i < 24; i++) {
        this.hourMap[i] = [];
      }
      for (let i = 0; i < this.list.length; i++) {
        const start = new Date(this.list[i].start);
        const stop = new Date(this.list[i].stop);
        const timeZoneOffset = (new Date().getTimezoneOffset() / 60) * -1;
        const absTime =
          (start.getTime() / 1000 -
            new Date(
              this.list[i].start.split('T')[0] +
                'T00:00:00' +
                (timeZoneOffset > 0
                  ? `+${('00' + timeZoneOffset).slice(-2)}`
                  : `${('00' + timeZoneOffset).slice(-2)}`) +
                ':00',
            ).getTime() /
              1000) /
          3600;
        let duration = (stop.getTime() / 1000 - start.getTime() / 1000) / 3600;

        let hour = ~~absTime;
        let offset = absTime - hour;
        let hourOffset = 0;
        let color = this.color[this.list[i].name];

        for (let j = 0; j < 10; j++) {
          let reminder = duration - (1 - offset);
          if (reminder > 0) {
            if (hour + hourOffset < 24) {
              this.hourMap[hour + hourOffset].push({ color, offset, duration: 1 - offset });
            }

            duration -= 1 - offset;
            hourOffset += 1;
            offset = 0;
          } else {
            if (hour + hourOffset < 24) {
              this.hourMap[hour + hourOffset].push({ color, offset, duration });
            }
            break;
          }
        }
      }

      this.isToday = Moment(this.date).format('YYYY-MM-DD') === Moment().format('YYYY-MM-DD');
    },
    async remove(item: any) {
      if (confirm('Are you sure you want to delete it?')) {
        await RestApi.task.delete(item.id);
        // @ts-ignore
        this.$root.refresh();
      }
    },
  },
  data: () => {
    return {
      isShowAddForm: false,
      isShowEditForm: false,
      isToday: false,
      itemId: '',

      hourMap: [] as any[],
      list: [] as any[],

      color: {
        /*programming: '#b404ca',
        draw: '#af2222',
        language: '#e68a02',*/
      } as any,
    };
  },
});
</script>

<style lang="scss" scoped>
.history {
  display: flex;
  height: calc(100%);

  .hour_map {
    margin-right: 10px;
    width: 180px;
    overflow-y: auto;
    height: max-content;
    max-height: 100%;

    .time {
      color: #999999;

      margin-bottom: 5px;
      font-size: 14px;
      display: flex;
      align-items: center;

      &:last-child {
        margin-bottom: 0;
      }

      .hour {
        margin-right: 5px;
      }

      .bar {
        height: 10px;
        background: #0e0e0e;
        width: 100%;
        position: relative;

        .fill {
          position: absolute;
          box-sizing: border-box;
          display: flex;
          align-items: center;
          height: 100%;
          background: #414141;
          font-weight: bold;
        }

        .current {
          position: absolute;
          top: -4px;
          color: #fe0000;
          font-weight: bold;
        }
      }
    }
  }

  .task_list {
    margin-right: 10px;
    width: 280px;
    overflow-y: auto;
    height: max-content;
    max-height: 100%;
  }
}
</style>

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
        </div>
      </div>
    </div>

    <div class="block description">
      <div class="header">
        Descriprion
        <img @click="isShowAddForm = true" class="clickable" src="../../asset/add.svg" alt="" />
      </div>
      <div class="body" style="color: #999999; line-height: 26px">
        <div v-for="x in list" :key="x.id" style="position: relative">
          <div><b>Start:</b> {{ $root.moment(x.start).format('YYYY-MM-DD HH:mm') }}</div>
          <div>
            <b>Duration:</b>
            {{
              $root
                .moment()
                .startOf('day')
                .seconds($root.moment(x.stop).diff($root.moment(x.start), 'seconds'))
                .format('HH:mm')
            }}
          </div>
          <div style="display: flex; position: absolute; right: 0; top: 0">
            <img
              @click="(isShowEditForm = true), (itemId = x.id)"
              class="clickable"
              src="../../asset/pencil.svg"
              alt=""
              style="margin-left: auto"
            />
            <img
              @click="remove(x)"
              class="clickable"
              src="../../asset/trash.svg"
              alt=""
              style="margin-left: 10px"
            />
          </div>

          <div style="margin-top: 20px" v-html="x.description.replace(/\n/g, '<br>')"></div>
        </div>
      </div>
    </div>

    <Add :date="date" v-if="isShowAddForm" @close="(isShowAddForm = false), refresh()" />
    <Edit :id="itemId" v-if="isShowEditForm" @close="(isShowEditForm = false), refresh()" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';
import Add from './Add.vue';
import Edit from './Edit.vue';
import Moment from 'moment';

export default defineComponent({
  props: {
    date: Object,
  },
  components: { Add, Edit },
  async mounted() {
    this.refresh();
  },
  watch: {
    date(value: Date) {
      this.refresh();
    },
  },
  methods: {
    async refresh() {
      this.list = await RestApi.sleep.getByDay(Moment(this.date).format('YYYY-MM-DD'));
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
        let color = '#666666';

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
    },

    async remove(item: any) {
      if (confirm('Are you sure you want to delete it?')) {
        await RestApi.sleep.delete(item.id);
        this.refresh();
      }
    },
  },
  data: () => {
    return {
      hourMap: [] as any[],
      list: [] as any[],
      isShowAddForm: false,
      isShowEditForm: false,
      itemId: '',
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
      }
    }
  }

  .description {
    margin-right: 10px;
    width: 360px;
    overflow-y: auto;
    height: max-content;
    max-height: 100%;
  }
}
</style>

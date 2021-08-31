<template>
  <div class="history">
    <div class="block hour_map">
      <!-- Header -->
      <div class="header">
        Hour Map
        <img @click="isShowAddForm = true" class="clickable" src="../../asset/add.svg" alt="" />
      </div>

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

    <Add :date="date" v-if="isShowAddForm" @close="(isShowAddForm = false), refresh()" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';
import Add from './Add.vue';
import Task from './Task.vue';
import Moment from 'moment';

export default defineComponent({
  props: {
    date: Object,
  },
  components: { Add, Task },
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
        const absTime =
          (start.getTime() / 1000 -
            new Date(this.list[i].start.split('T')[0] + ' 00:00:00').getTime() / 1000) /
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
    },
  },
  data: () => {
    return {
      hourMap: [] as any[],
      list: [] as any[],
      isShowAddForm: false,
      color: {
        programming: '#b404ca',
        draw: '#af2222',
        language: '#e68a02',
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
    width: 220px;
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
        height: 20px;
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
}
</style>

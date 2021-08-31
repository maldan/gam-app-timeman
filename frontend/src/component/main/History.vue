<template>
  <div class="history">
    <div class="block hour_map">
      <!-- Header -->
      <div class="header">Hour Map</div>

      <div class="time" v-for="(x, i) in hourMap" :key="x">
        <div class="hour">{{ ('0' + i).slice(-2) }}</div>
        <div class="bar">
          <div class="fill" :style="{ width: x / 60 + '%' }"></div>
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
        <Task :item="x" :nextItem="list[i + 1]" :date="date" />
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
      this.list = await RestApi.task.getByDay(Moment(this.date).format('YYYY-MM-DD'));
      for (let i = 0; i < 24; i++) {
        this.hourMap[i] = 0;
      }
    },
  },
  data: () => {
    return {
      hourMap: [] as number[],
      list: [] as any[],
      isShowAddForm: false,
    };
  },
});
</script>

<style lang="scss" scoped>
.history {
  display: flex;
  height: calc(100% - 10px);

  .hour_map {
    margin-right: 10px;
    width: 160px;

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

        .fill {
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

  .task_list {
    margin-right: 10px;
    width: 260px;
  }
}
</style>

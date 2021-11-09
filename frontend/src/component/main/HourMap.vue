<template>
  <div :class="$style.hour_map">
    <div :class="$style.time" v-for="(list, i) in $store.state.task.hourMap" :key="i">
      <div :class="$style.hour">{{ ('0' + i).slice(-2) }}</div>
      <div :class="$style.bar">
        <div
          v-for="x in list"
          :key="x.offset"
          :class="[$style.fill, $style.clickable]"
          :style="{
            background: x.color,
            left: x.offset * 100 + '%',
            width: x.duration * 100 + '%',
          }"
        ></div>
        <div
          :class="$style.current"
          v-if="$store.state.task.isToday && new Date().getHours() === i"
          :style="{ left: (new Date().getMinutes() / 60) * 100 + '%' }"
        >
          |
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {},
  components: {},
  async mounted() {},
  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" module>
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
</style>

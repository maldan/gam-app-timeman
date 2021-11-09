<template>
  <div :class="$style.main">
    <ui-block title="hour map">
      <HourMap />
    </ui-block>

    <ui-block
      :class="$style.block"
      title="history"
      icon="plus"
      :scrollY="true"
      @iconClick="
        $store.dispatch('modal/show', {
          name: 'add/task',
          data: {
            name: '',
            description: '',
            start: $root.moment($store.state.task.date).format('YYYY-MM-DD HH:mm:ss'),
            stop: $root.moment($store.state.task.date).format('YYYY-MM-DD HH:mm:ss'),
          },
          onSuccess: () => {
            $store.dispatch('task/add');
          },
        })
      "
    >
      <div class="task" v-for="(x, i) in $store.state.task.history" :key="x">
        <Task
          :item="x"
          :nextItem="$store.state.task.history[i + 1]"
          :date="$store.state.task.date"
        />
      </div>
    </ui-block>

    <ui-block title="schedule">
      <ui-schedule
        :map="$store.state.task.yearMap"
        :max="28800"
        :date="$store.state.task.date"
        :yearRange="[-4, -3, -2, -1, 0]"
        formatValue="time"
        color="#2aefca"
        @select="$store.dispatch('task/setDate', $event)"
      />
    </ui-block>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import HourMap from '../component/main/HourMap.vue';
import Task from '../component/main/Task.vue';

export default defineComponent({
  components: { HourMap, Task },
  async mounted() {},
  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style module lang="scss">
@import '../gam_sdk_ui/vue/style/size.scss';
@import '../gam_sdk_ui/vue/style/color.scss';

.main {
  box-sizing: border-box;
  display: grid;
  height: calc(100% - 48px);
  gap: $gap-base;
  grid-template-columns: 200px 340px 1fr;
  padding: $gap-base;

  .block {
    min-height: 0;
  }
}
</style>

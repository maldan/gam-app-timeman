<template>
  <div class="main">
    <Header />

    <div class="main_body">
      <History :date="currentDate" @update="description = $event" />
      <Schedule @select="(currentDate = $event), refresh()" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from '../component/Header.vue';
import History from '../component/sleep/History.vue';
import Schedule from '../component/sleep/Schedule.vue';

export default defineComponent({
  components: { Header, History, Schedule },
  async mounted() {
    this.refresh();
  },
  methods: {
    async refresh() {},
  },
  data: () => {
    return {
      currentDate: new Date(),
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  box-sizing: border-box;
  height: 100%;
  padding: 10px;

  .main_body {
    margin-top: 10px;
    display: flex;
    height: calc(100% - 60px);
  }
}
</style>

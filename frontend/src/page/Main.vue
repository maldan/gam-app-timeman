<template>
  <div class="main">
    <Header />

    <div class="body">
      <History :date="currentDate" />
      <Schedule @select="currentDate = $event" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from '../component/Header.vue';
import History from '../component/main/History.vue';
import Schedule from '../component/main/Schedule.vue';

export default defineComponent({
  components: { Header, History, Schedule },
  async mounted() {},
  methods: {},
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

  .body {
    margin-top: 10px;
    display: flex;
    height: calc(100% - 60px);
  }
}
</style>

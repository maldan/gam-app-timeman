<template>
  <div class="main">
    <Header />

    <div class="main_body">
      <History :date="currentDate" @update="description = $event" />
      <div style="flex: 1">
        <Schedule @select="(currentDate = $event), refresh()" />
        <div class="block" style="margin-top: 10px">
          <div class="header">Descriprion</div>
          <div
            class="body"
            style="color: #999999; line-height: 26px"
            v-html="description.join('<br><br>')"
          ></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from '../component/Header.vue';
import History from '../component/sleep/History.vue';
import Schedule from '../component/sleep/Schedule.vue';
import { RestApi } from '../util/RestApi';
import Moment from 'moment';

export default defineComponent({
  components: { Header, History, Schedule },
  async mounted() {
    this.refresh();
  },
  methods: {
    async refresh() {
      // @ts-ignore
      const list = await RestApi.sleep.getByDay(Moment(this.currentDate).format('YYYY-MM-DD'));
      this.description.length = 0;
      for (let i = 0; i < list.length; i++) {
        if (list[i].description) {
          this.description.push(list[i].description.replace(/\n/g, '<br>'));
        }
      }
    },
  },
  data: () => {
    return {
      currentDate: new Date(),
      description: [] as any[],
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

<template>
  <router-view />
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Moment from 'moment';

export default defineComponent({
  components: {},
  async mounted() {},
  methods: {
    refresh() {
      this.refreshId = Math.random();
    },
  },
  data: () => {
    return {
      // @ts-ignore
      API_URL: process.env.VUE_APP_API_URL || `${window.location.origin}/api`,
      moment: Moment,
      refreshId: 0,
    };
  },
});
</script>

<style lang="scss" scoped></style>

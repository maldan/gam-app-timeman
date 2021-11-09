<template>
  <ui-header :items="header" />
  <router-view />
  <ui-modal-view />
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Moment from 'moment';

export default defineComponent({
  components: {},
  async mounted() {},
  methods: {},
  data: () => {
    return {
      // @ts-ignore
      API_URL: process.env.VUE_APP_API_URL || `${window.location.origin}/api`,
      moment: Moment,
      console: console,
      header: [
        { name: 'Main', url: '/' },
        { name: 'Sleep', url: '/sleep' },
      ],
    };
  },
});
</script>

<style lang="scss" scoped></style>

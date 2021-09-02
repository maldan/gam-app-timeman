<template>
  <div :class="$style.container">
    <div :class="$style.window">
      <Input icon="food" placeholder="Name..." style="margin-bottom: 10px" v-model="name" />
      <TextArea
        icon="weight"
        placeholder="Description..."
        style="margin-bottom: 10px"
        v-model="description"
      />

      <Input
        icon="date"
        type="datetime-local"
        placeholder="Start..."
        style="margin-bottom: 10px"
        v-model="start"
      />
      <Input
        icon="date"
        type="datetime-local"
        placeholder="Stop..."
        style="margin-bottom: 10px"
        v-model="stop"
      />

      <div style="display: flex">
        <Button @click="$emit('close')" text="Cancel" style="margin-right: 5px" />
        <Button @click="submit()" text="Add" icon="add" style="margin-left: 5px" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';
import Button from '../Button.vue';
import Input from '../Input.vue';
import Select from '../Select.vue';
import TextArea from '../TextArea.vue';

export default defineComponent({
  props: {
    date: Object,
  },
  components: { Button, Input, Select, TextArea },
  async mounted() {},
  methods: {
    async submit() {
      await RestApi.task.add(this.name, this.description, this.start, this.stop);
      this.$emit('close');
    },
  },
  data() {
    const moment = (this.$root as any)['moment'];
    const d = this.date as Date;
    d.setHours(new Date().getHours());
    d.setMinutes(new Date().getMinutes());
    d.setSeconds(new Date().getSeconds());

    return {
      name: '',
      description: '',
      start: moment(d).format('YYYY-MM-DDTHH:mm:ss'),
      stop: moment(d).format('YYYY-MM-DDTHH:mm:ss'),
    };
  },
});
</script>

<style lang="scss" module>
.container {
  position: fixed;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.6);
  z-index: 10;
  display: flex;
  align-items: center;
  justify-content: center;

  .window {
    min-width: 320px;
    background: #313131a8;
    border-radius: 4px;
    padding: 25px;
    color: #fefefe;
    box-shadow: 0px 1px 6px 2px #00000055;

    .found_item {
      padding: 10px 15px;
      background: #3a3a3a;
      margin-bottom: 10px;
      border-radius: 4px;
      display: flex;

      img {
        margin-left: auto;
      }
    }
  }
}
</style>

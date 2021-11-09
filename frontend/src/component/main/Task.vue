<template>
  <div :class="$style.block">
    <div :class="$style.header">
      <div :class="$style.left">
        {{ item.name }}
      </div>
      <div :class="$style.right">{{ $root.moment(item.start).format('HH:mm') }}</div>
      <div :class="$style.right2">
        {{
          $root
            .moment()
            .startOf('day')
            .seconds($root.moment(item.stop).diff($root.moment(item.start), 'seconds'))
            .format('HH:mm')
        }}
      </div>
    </div>
    <div :class="$style.body">
      <div :class="$style.name">
        <div v-html="item.description.replace(/\n/g, '<br>')"></div>

        <div :class="$style.icons">
          <ui-button
            @click="
              $store.dispatch('modal/show', {
                name: 'add/task',
                data: {
                  ...item,
                  start: $root.moment(item.start).format('YYYY-MM-DD HH:mm:ss'),
                  stop: $root.moment(item.stop).format('YYYY-MM-DD HH:mm:ss'),
                },
                onSuccess: () => {
                  $store.dispatch('task/add');
                },
              })
            "
            :class="$style.button"
            icon="copy"
            size="compact-square"
          />
          <ui-button
            @click="
              $store.dispatch('modal/show', {
                name: 'edit/task',
                data: {
                  ...item,
                },
                onSuccess: () => {
                  $store.dispatch('task/update');
                },
              })
            "
            :class="$style.button"
            icon="pencil"
            size="compact-square"
          />
          <ui-button
            @click="
              $store.dispatch('modal/show', {
                name: 'approve',
                data: {
                  title: 'Remove this task?',
                },
                onSuccess: () => {
                  $store.dispatch('task/remove', item.id);
                },
              })
            "
            :class="$style.button"
            icon="trash"
            size="compact-square"
          />
        </div>
      </div>
    </div>
    <div :class="$style.gap" v-if="gap()">{{ gap() }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Moment from 'moment';

export default defineComponent({
  props: {
    item: Object,
    nextItem: Object,
    date: Object,
  },
  components: {},
  async mounted() {},
  methods: {
    gap() {
      if (!this.item) {
        return null;
      }
      if (
        !this.nextItem &&
        Moment(this.date).format('YYYY-MM-DD') !== Moment(new Date()).format('YYYY-MM-DD')
      ) {
        return null;
      }

      const a = new Date(this.item.stop).getTime() / 1000;
      const b = new Date(this.nextItem?.start || new Date()).getTime() / 1000;
      if (b - a < 60 * 5) {
        return null;
      }
      return Moment.utc(Moment.duration(b - a, 'seconds').asMilliseconds()).format('HH:mm');
    },
  },
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" module>
@import '../../gam_sdk_ui/vue/style/size.scss';
@import '../../gam_sdk_ui/vue/style/color.scss';

.block {
  font-size: 15px;
  margin-bottom: 5px;

  .header {
    display: flex;

    .left,
    .right,
    .right2 {
      padding: 5px 10px;
      border-radius: 6px 6px 0 0;
      color: $text-gray;
      background: $gray-medium;
      font-weight: bold;

      span {
        color: $text-gray;
        font-style: italic;
        font-weight: 300;
      }
    }

    .left {
      color: $text-green;
    }

    .right {
      margin-left: auto;
    }

    .right2 {
      margin-left: $gap-base;
      color: $text-green;
    }
  }

  .body {
    padding: $gap-base;
    background: $gray-medium;
    border-radius: 0 0 6px 6px;
    color: $text-gray;

    .name {
      display: flex;
      color: $text-gray;

      .icons {
        margin-left: auto;
        display: flex;

        .button {
          margin-left: $gap-base;
          padding: 2px;
        }
      }
    }

    .component {
      margin-top: 10px;
      display: grid;
      align-items: center;
      grid-template-columns: 1fr 1fr 1fr;
      gap: 5px;
      font-size: 14px;

      > div {
        display: flex;

        span {
          color: #828282;
          margin-right: auto;
        }
        flex: 1;
        background: #80808045;
        padding: 5px 10px;
      }
    }
  }

  .gap {
    background: $gray-medium;
    padding: 3px 10px;
    text-align: center;
    border-radius: 4px;
    margin-top: 25px;
    margin-bottom: 25px;
    color: $text-green;
    position: relative;
    width: max-content;
    margin-left: auto;
    margin-right: auto;

    &::before {
      content: '|';
      position: absolute;
      left: 50%;
      top: -22px;
      color: $text-gray;
    }

    &::after {
      content: '|';
      position: absolute;
      left: 50%;
      bottom: -20px;
      color: $text-gray;
    }
  }
}
</style>

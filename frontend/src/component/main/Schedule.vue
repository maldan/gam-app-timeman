<template>
  <div class="block" style="flex: 1; height: max-content">
    <div class="header">Schedule</div>
    <div class="body">
      <div :class="$style.schedule">
        <div :class="$style.year">
          <div :class="$style.month" v-for="(x, i) in month" :key="x">
            <div>{{ x }}</div>
            <div style="display: grid; margin-top: 10px; grid-template-columns: repeat(7, 15px)">
              <div
                @click="$emit('select', z), (date = z)"
                class="clickable"
                :class="[
                  $style.cell,
                  $root.moment(new Date()).format('DD MMM YYYY') ===
                  $root.moment(z).format('DD MMM YYYY')
                    ? $style.today
                    : '',
                  $root.moment(date).format('DD MMM YYYY') === $root.moment(z).format('DD MMM YYYY')
                    ? $style.selected
                    : '',
                  z === null ? $style.null : '',
                ]"
                :style="{
                  background: getPower(map[$root.moment(z).format('YYYY-MM-DD')]),
                }"
                v-for="z in days[i]"
                :key="z"
                :title="(map[$root.moment(z).format('YYYY-MM-DD')] / 3600).toFixed(2)"
              ></div>
            </div>
          </div>
        </div>

        <div style="margin-top: 15px; color: #979797; display: flex; align-items: center">
          <div style="margin-right: 15px">
            Date: {{ $root.moment(date).format('DD MMM YYYY') }} | Total:
            {{ totalTime(total) }}
          </div>

          <button
            v-for="x in [-4, -3, -2, -1, 0]"
            :class="[
              $style.change_year,
              date.getFullYear() === new Date().getFullYear() + x ? $style.selected : '',
            ]"
            class="clickable"
            :key="x"
            @click="
              (date = new Date(
                new Date().getFullYear() + x,
                new Date().getMonth(),
                new Date().getDate(),
              )),
                $emit('select', date),
                refresh()
            "
          >
            {{ new Date().getFullYear() + x }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';
import Moment from 'moment';

export default defineComponent({
  props: {
    item: Object,
    type: String,
    mode: String,
  },
  components: {},
  async mounted() {
    this.refresh();

    this.keyboardEvent = document.addEventListener('keydown', (e: KeyboardEvent) => {
      if (e.key === 'ArrowRight') {
        // @ts-ignore
        this.date = Moment(this.date).add(1, 'day').toDate();
        this.$emit('select', this.date);
      }
      if (e.key === 'ArrowLeft') {
        // @ts-ignore
        this.date = Moment(this.date).add(-1, 'day').toDate();
        this.$emit('select', this.date);
      }
      if (e.key === 'ArrowUp') {
        // @ts-ignore
        this.date = Moment(this.date).add(-7, 'day').toDate();
        this.$emit('select', this.date);
      }
      if (e.key === 'ArrowDown') {
        // @ts-ignore
        this.date = Moment(this.date).add(7, 'day').toDate();
        this.$emit('select', this.date);
      }
    });
  },
  beforeUnmount() {
    document.removeEventListener('keydown', this.keyboardEvent);
  },
  methods: {
    totalTime(time: number) {
      const h = ~~(time / 3600);
      const m = (time / 60) % 60;
      return `${h} h ${~~m} m`;
    },
    getPower(power: number) {
      if (power === undefined) return 'transparent';
      if (power <= 0) return '#e2e2e223';

      let p = power / 3600 / 6;
      if (p < 0.3) {
        p = 0.3;
      }
      return `rgba(39, 255, 0, ${p})`;
    },
    async refresh() {
      this.days.length = 0;

      for (let i = 0; i < 12; i++) {
        const l = this.getDates(i);
        const ll = [];
        let dd = l[0].getDay() - 1;
        if (dd < 0) {
          dd = 6;
        }
        for (let j = 0; j < dd; j++) {
          ll.push(null);
        }
        this.days[i] = [...ll, ...l];
      }

      // @ts-ignore
      this.map = await RestApi.task.getYearMap(Moment(this.date).format('YYYY-MM-DD'));
      this.total = 0;
      for (let x in this.map) {
        this.total += ~~(this.map as any)[x] as number;
      }
    },
    getDates(month: number): Date[] {
      const year = this.date.getFullYear();
      const out = [];
      for (let i = 1; i <= 32; i++) {
        const cFrom = new Date(
          `${year}-${('00' + (month + 1)).slice(-2)}-${('00' + i).slice(-2)}T00:00:00.000+00:00`,
        );
        if (cFrom.toString() === 'Invalid Date') {
          continue;
        }
        if (cFrom.getMonth() !== month) {
          break;
        }

        out.push(cFrom);
      }
      return out;
    },
  },
  data: () => {
    return {
      month: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
      days: [] as any[],
      date: new Date(),
      map: {},
      total: 0,
      keyboardEvent: null as any,
    };
  },
});
</script>

<style lang="scss" module>
.schedule {
  display: flex;
  flex-direction: column;
  box-sizing: border-box;

  .year {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    gap: 10px;

    .month {
      flex: 1;
      font-size: 12px;
      color: #acacac;

      .cell {
        width: 12px;
        height: 12px;
        background: #2b2b2b;
        font-size: 8px;
        display: flex;
        align-items: center;
        justify-content: center;
        margin-bottom: 2px;

        &.null {
          pointer-events: none;
          background: transparent;
        }

        &.selected {
          border: 1px solid #ca6f00;
        }

        &.today::before {
          content: '●';
          display: flex;
          align-items: center;
          justify-content: center;
          color: #fefefe;
        }
      }
    }
  }

  .change_year {
    background: #525252;
    padding: 5px 15px;
    border: 0;
    border-radius: 3px;
    color: #fefefe;
    font-weight: bold;
    color: #c5c5c5;
    margin-right: 5px;

    &.selected {
      background: #0075ce;
      color: #fefefe;
    }
  }
}
</style>

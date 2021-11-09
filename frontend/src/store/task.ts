import Axios from 'axios';
import Moment from 'moment';
import { ActionContext } from 'vuex';
import { MainTree } from '.';

export type TaskStore = {
  history: any[];
  yearMap: any[];
  date: Date;
  hourMap: any[];
  isToday: boolean;
};
export type TaskActionContext = ActionContext<TaskStore, MainTree>;

export default {
  namespaced: true,
  state: {
    history: [],
    yearMap: {},
    date: new Date('1970-01-01'),
    hourMap: [],
    isToday: false,
  },
  mutations: {
    SET_HISTORY(state: TaskStore, history: any[]) {
      state.history = history;
    },
    SET_YEAR_MAP(state: TaskStore, payload: any) {
      state.yearMap = payload;
    },
    SET_HOUR_MAP(state: TaskStore, payload: any) {
      state.hourMap = payload;
    },
    SET_DATE(state: TaskStore, date: Date) {
      const d = new Date(date);
      d.setHours(new Date().getHours());
      d.setMinutes(new Date().getMinutes());
      d.setSeconds(new Date().getSeconds());
      state.date = d;
      state.isToday = Moment(d).format('YYYY-MM-DD') === Moment().format('YYYY-MM-DD');
    },
  },
  actions: {
    async getHistory(action: TaskActionContext) {
      const list = (
        await Axios.get(
          `${action.rootState.main.API_URL}/task/byDay?date=${Moment(action.state.date).format(
            'YYYY-MM-DD',
          )}`,
        )
      ).data.response;
      action.commit('SET_HISTORY', list);

      action.dispatch('calculateHourMap');
    },
    async calculateHourMap(action: TaskActionContext) {
      const hourMap = [] as any[];
      for (let i = 0; i < 24; i++) {
        hourMap[i] = [];
      }
      for (let i = 0; i < action.state.history.length; i++) {
        const start = new Date(action.state.history[i].start);
        const stop = new Date(action.state.history[i].stop);
        const timeZoneOffset = (new Date().getTimezoneOffset() / 60) * -1;
        const absTime =
          (start.getTime() / 1000 -
            new Date(
              action.state.history[i].start.split('T')[0] +
                'T00:00:00' +
                (timeZoneOffset > 0
                  ? `+${('00' + timeZoneOffset).slice(-2)}`
                  : `${('00' + timeZoneOffset).slice(-2)}`) +
                ':00',
            ).getTime() /
              1000) /
          3600;
        let duration = (stop.getTime() / 1000 - start.getTime() / 1000) / 3600;

        const hour = ~~absTime;
        let offset = absTime - hour;
        let hourOffset = 0;
        const color = action.rootState.activity.color[action.state.history[i].name];

        for (let j = 0; j < 10; j++) {
          const reminder = duration - (1 - offset);
          if (reminder > 0) {
            if (hour + hourOffset < 24) {
              hourMap[hour + hourOffset].push({ color, offset, duration: 1 - offset });
            }

            duration -= 1 - offset;
            hourOffset += 1;
            offset = 0;
          } else {
            if (hour + hourOffset < 24) {
              hourMap[hour + hourOffset].push({ color, offset, duration });
            }
            break;
          }
        }
      }

      action.commit('SET_HOUR_MAP', hourMap);
    },
    async getYearMap(action: TaskActionContext) {
      const yearMap = (
        await Axios.get(
          `${action.rootState.main.API_URL}/task/yearMap?date=${Moment(action.state.date).format(
            'YYYY-MM-DD',
          )}`,
        )
      ).data.response;
      action.commit('SET_YEAR_MAP', yearMap);
    },
    setDate(action: TaskActionContext, date: Date) {
      const isRefresh = date.getFullYear() != action.state.date.getFullYear();
      action.commit('SET_DATE', date);
      if (isRefresh) {
        action.dispatch('getYearMap');
      }
      action.dispatch('getHistory');
    },
    async add(action: TaskActionContext) {
      await Axios.post(`${action.rootState.main.API_URL}/task`, action.rootState.modal.data);

      await action.dispatch('getHistory');
      await action.dispatch('getYearMap');
    },
    async update(action: TaskActionContext) {
      await Axios.patch(`${action.rootState.main.API_URL}/task`, action.rootState.modal.data);

      await action.dispatch('getHistory');
      await action.dispatch('getYearMap');
    },
    async remove(action: TaskActionContext, taskId: string) {
      await Axios.delete(`${action.rootState.main.API_URL}/task?id=${taskId}`);

      await action.dispatch('getHistory');
      await action.dispatch('getYearMap');
    },
  },
};

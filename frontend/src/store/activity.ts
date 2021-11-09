import Axios from 'axios';
import Moment from 'moment';
import { ActionContext } from 'vuex';
import { MainTree } from '.';

export type ActivityStore = {
  list: any[];
  color: any;
};
export type ActivityActionContext = ActionContext<ActivityStore, MainTree>;

export default {
  namespaced: true,
  state: {
    list: [],
    color: {},
    hourRate: {},
  },
  mutations: {
    SET_LIST(state: ActivityStore, list: any[]) {
      state.list = list;

      for (let i = 0; i < list.length; i++) {
        state.color[list[i].name] = list[i].color || '#cccccc';
      }
    },
  },
  actions: {
    async getList(action: ActivityActionContext) {
      const list = (await Axios.get(`${action.rootState.main.API_URL}/activity/list`)).data
        .response;
      action.commit('SET_LIST', list);
    },
  },
};

import { createStore } from 'vuex';

import modal, { ModalStore } from '../gam_sdk_ui/vue/store/modal';
import main, { MainStore } from './main';
import task, { TaskStore } from './task';
import activity, { ActivityStore } from './activity';
import sleep, { SleepStore } from './sleep';

export type MainTree = {
  main: MainStore;
  modal: ModalStore;
  task: TaskStore;
  activity: ActivityStore;
  sleep: SleepStore;
};

export default createStore({
  modules: { main, modal, task, activity, sleep },
});

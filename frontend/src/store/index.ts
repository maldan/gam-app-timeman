import { createStore } from 'vuex';

import modal, { ModalStore } from '../gam_sdk_ui/vue/store/modal';
import main, { MainStore } from './main';
import task, { TaskStore } from './task';
import activity, { ActivityStore } from './activity';

export type MainTree = {
  main: MainStore;
  modal: ModalStore;
  task: TaskStore;
  activity: ActivityStore;
};

export default createStore({
  modules: { main, modal, task, activity },
});

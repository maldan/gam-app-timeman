import Axios from 'axios';

const API_URL = process.env.VUE_APP_API_URL || `${window.location.origin}/api`;

export const RestApi = {
  task: {
    async add(name: string, description: string, start: string, stop: string) {
      return (
        await Axios.post(`${API_URL}/task`, {
          name,
          description,
          start,
          stop,
        })
      ).data.response;
    },

    async getByDay(date: string): Promise<any> {
      return (await Axios.get(`${API_URL}/task/byDay?date=${date}`)).data.response;
    },

    async getYearMap(date: string): Promise<any> {
      return (await Axios.get(`${API_URL}/task/yearMap?date=${date}`)).data.response;
    },
  },

  sleep: {
    async add(description: string, start: string, stop: string) {
      return (
        await Axios.post(`${API_URL}/sleep`, {
          description,
          start,
          stop,
        })
      ).data.response;
    },

    async getByDay(date: string): Promise<any[]> {
      return (await Axios.get(`${API_URL}/sleep/byDay?date=${date}`)).data.response;
    },

    async getYearMap(date: string): Promise<any> {
      return (await Axios.get(`${API_URL}/sleep/yearMap?date=${date}`)).data.response;
    },
  },
};

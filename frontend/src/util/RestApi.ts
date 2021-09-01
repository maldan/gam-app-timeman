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
    async update(id: string, name: string, description: string, start: string, stop: string) {
      return (
        await Axios.patch(`${API_URL}/task`, {
          id,
          name,
          description,
          start,
          stop,
        })
      ).data.response;
    },

    async get(id: string) {
      return (await Axios.get(`${API_URL}/task?id=${id}`)).data.response;
    },

    async delete(id: string) {
      return (await Axios.delete(`${API_URL}/task?id=${id}`)).data.response;
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

    async update(id: string, description: string, start: string, stop: string) {
      return (
        await Axios.patch(`${API_URL}/sleep`, {
          id,
          description,
          start,
          stop,
        })
      ).data.response;
    },

    async get(id: string) {
      return (await Axios.get(`${API_URL}/sleep?id=${id}`)).data.response;
    },

    async delete(id: string) {
      return (await Axios.delete(`${API_URL}/sleep?id=${id}`)).data.response;
    },

    async getByDay(date: string): Promise<any[]> {
      return (await Axios.get(`${API_URL}/sleep/byDay?date=${date}`)).data.response;
    },

    async getYearMap(date: string): Promise<any> {
      return (await Axios.get(`${API_URL}/sleep/yearMap?date=${date}`)).data.response;
    },
  },
};

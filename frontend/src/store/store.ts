import { defineStore } from 'pinia'
import { SayHello } from '../../wailsjs/go/backend/Config'

export const configStore = defineStore({
  id: 'config',
  state: () => {
    return {
      name: 'hello',
    }
  },
  getters: {
    async hello() {
      let res = await SayHello()
      return res
    },
  },
  actions: {
    async helloAction(){
      return await  SayHello()
    }
  },
})

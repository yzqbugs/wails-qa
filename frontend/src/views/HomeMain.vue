<script lang="ts" setup>
import { reactive } from 'vue'
import { SayHello } from '../../wailsjs/go/backend/Config'
import { configStore } from '@/store/store'
let store = configStore()
const data = reactive({
  name: '',
  resultText: 'Please enter your name below 👇',
})
async function setconf() {
  let res = await SayHello()
  if (res) {
    console.log(`%csetconf success`, `color:red;font-size:16px;background:transparent`)
  }
  data.resultText = res
}
async function getPiniaHello(){
  data.resultText=await  store.helloAction()
}
async function setPinia() {
  if (await store.hello) {
    console.log(`%cpinia success`, `color:red;font-size:16px;background:transparent`)
  }
  data.resultText = await store.hello
}
</script>

<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>

    <button class="btn" @click="setconf">setconf</button>
    <button class="btn" @click="setPinia">setPinia</button>
    <button class="btn" @click="getPiniaHello">getPiniaHello</button>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>

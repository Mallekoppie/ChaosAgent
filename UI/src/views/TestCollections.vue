<template>
  <div class="custom-layout-grouping">
    <h1 class="text-white m-2">Group: {{testGroupName}}</h1>
    <p class="text-gray-400 m-2">{{testGroupDescription}}</p>
    <table class="table-auto border-2 rounded-t-lg text-gray-400 hover stripe m-2">
      <tr class="text-left border-b border-gray-300">
        <th class="px-4 py-3">Name</th>
        <th class="px-4 py-3">Description</th>
        <th class="px-4 py-3">Open</th>
      </tr>
      <tr v-for="item in testCollections" v-bind:key="item.id" class="bg-gray-700 border-b border-gray-600 hover:bg-black">
        <td class="px-4 py-3">{{ item.name }}</td>
        <td class="px-4 py-3">{{ item.description }}</td>
        <td class="px-4 py-3"><button @click="openTestCollection(item)" class="custom-button-normal">Open</button></td>
      </tr>
    </table>
  </div>
</template>

<script>
import {dataStore} from '../shared/datastoretemp'
import router from "@/router";

export default {
  name: "TestCollections",
  props:{
    testGroupInput: Object
  },
  data() {
    return {
      testCollections: [],
      testGroupName: '',
      testGroupDescription: ''
    }
  },created() {
    console.log('Creating new instance. Logging testGroupinput')
    console.log(this.testGroupInput)
    console.log(this.testGroupInput.name)
    let input = this.testGroupInput
    this.testGroupName = input.name
    this.testGroupDescription = input.description
    this.testCollections = dataStore.getTestCollections()
  },
  methods:{
    openTestCollection(id){
      console.log(id)
      router.push({name: 'TestCollection', params: id})
    }
  }
}
</script>
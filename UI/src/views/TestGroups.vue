<template>
  <div class="custom-layout-grouping">
    <h1 class="text-white m-2">Test Groups</h1>
    <table class="table-auto border-2 rounded-t-lg text-gray-400 hover stripe">
      <tr class="text-left border-b border-gray-300">
        <th class="px-4 py-3">Name</th>
        <th class="px-4 py-3">Description</th>
        <th class="px-4 py-3">Open</th>
      </tr>
      <tr v-for="item in testGroups" v-bind:key="item.id" class="bg-gray-700 border-b border-gray-600 hover:bg-black">
        <td class="px-4 py-3">{{ item.name }}</td>
        <td class="px-4 py-3">{{ item.description }}</td>
        <td class="px-4 py-3">
            <button @click="openTestGroup(item)" class="custom-button-normal">Open</button>
        </td>
      </tr>
    </table>
  </div>
</template>

<script>
import {dataStore} from '../shared/datastoretemp'
import router from "@/router"

export default {
  name: "TestGroups",
  data() {
    return {
      testGroups: []
    }
  },
  created() {
    this.loadTestGroups()
  },
  methods: {
    loadTestGroups() {
      this.testGroups = dataStore.getTestGroups()
    },
    openTestGroup(input) {

      router.push({name: 'TestCollections', params: input})
    }
  }
}
</script>
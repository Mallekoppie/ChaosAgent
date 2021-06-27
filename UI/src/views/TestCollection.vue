<template>
  <div class="grid grid-cols-1">
    <div class="custom-layout-grouping m-2">
      <h1 class="text-white">
        {{ testCollection.name }}
      </h1>
      <label class="text-gray-400">
        {{ testCollection.description }}
      </label>
    </div>
    <div class="m-2">
      <h1 class="text-white">
        Tests
      </h1>
    </div>
    <div class="custom-layout-grouping grid grid-cols-1 m-2" v-for="test in testCollection.tests" v-bind:key="test.id">
      <div>
        <label>Test Name: </label>
        <input v-model="test.name">
      </div>

      <div>
        <label>Method: {{ test.method }}</label>

      </div>
      <h3 class="text-white">Request Headers</h3>
      <div v-for="header in test.headers" v-bind:key="header.id" class="m-2 p-2">
        <label class="m-2">Name</label>
        <input v-model="header.name" class="m-2">
        <label class="m-2">Value</label>
        <input v-model="header.value" class="m-2">
        <button class="custom-button-danger">Delete</button>
      </div>
      <label>Request Body</label>
      <textarea v-model="test.body" class="m-2 p-2"/>
      <label>Response Code</label>
      <input v-model="test.responseCode">
      <label>Response Body</label>
      <textarea v-model="test.responseBody"></textarea>
      <button class="custom-button-danger m-2">Delete Test</button>
    </div>
  </div>
</template>

<script>
import {dataStore} from "@/shared/datastoretemp";

export default {
  name: "TestCollection",
  props: {
    testCollectionInputId: Object
  },
  data() {
    return {
      testCollection: {}
    }
  },
  created() {
    console.log('Entered test collection id: ' + this.testCollectionInputId)
    this.testCollection = dataStore.getTestDetail()
  }
}
</script>
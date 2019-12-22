<template>
  <div>
    <b-list-group>
      <b-list-group-item v-for="testGroup in testGroups" :key="testGroup.id">
          <b-form>
              <b-form-group label="Name">
                  {{testGroup.name}}
              </b-form-group>
              <b-form-group label="Description">
                  {{testGroup.description}}
              </b-form-group>
          </b-form>                
        <b-button @click="showDetail(testGroup)">Show/Hide Test Collections</b-button>
        <b-list-group v-show="testGroup.showDetail">
          <b-list-group-item
            v-for="testCollection in testGroup.testCollections"
            :key="testCollection.id"
          >
            <span>Name: {{testCollection.name}}</span>
            <span>Description: {{testCollection.description}}</span>
          </b-list-group-item>
        </b-list-group>
      </b-list-group-item>
    </b-list-group>
  </div>
</template>

<script>
import { dataStore } from "@/shared/datastoretemp.js";

export default {
  name: "Tests",
  data() {
    return {
      testGroups: []
    };
  },
  created() {
    this.testGroups = dataStore.getTestGroups();
  },
  methods: {
    showDetail(item) {      
      if (item.showDetail == true) {
        item.showDetail = false;
      } else {
        item.showDetail = true;
      }
    }
  }
};
</script>

<template>
  <div>
    <b-list-group>
      <b-list-group-item v-for="testGroup in testGroups" :key="testGroup.id">
        <b-form>
          <b-form-group label="Name">{{ testGroup.name }}</b-form-group>
          <b-form-group label="Description">
            {{ testGroup.description }}
          </b-form-group>
        </b-form>
        <b-button @click="showDetail(testGroup)"
          >Show/Hide Test Collections</b-button
        >
        <b-list-group v-show="testGroup.showDetail">
          <b-list-group-item
            v-for="testCollection in testGroup.testCollections"
            :key="testCollection.id"
          >
            <b-row>
              <b-col lg="2">
                <span>Name: {{ testCollection.name }}</span>
              </b-col>
              <b-col lg="6">
                <span>Description: {{ testCollection.description }}</span>
              </b-col>
              <b-col lg="2">
                <router-link
                  tag="button"
                  class="btn btn-warning"
                  :to="{
                    name: 'test-detail',
                    params: { id: testCollection.id }
                  }"
                  >Edit</router-link
                >
              </b-col>
              <b-col lg="1">
                <b-button>Execute</b-button>
              </b-col>
            </b-row>
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

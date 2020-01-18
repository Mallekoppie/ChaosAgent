<template>
  <div>
    <b-form>
      <b-form-group label="Name">
        <b-form-input v-model="testGroup.name"> </b-form-input>
      </b-form-group>
      <b-form-group label="Description">
        <b-form-input v-model="testGroup.description"> </b-form-input>
      </b-form-group>
    </b-form>
    <b-form-row>
      <b-col cols="10">
        <b-button @click="showDetail(testGroup)"
          >Show/Hide Test Collections</b-button
        >
      </b-col>
      <b-col>
        <b-button variant="success">Save</b-button>
      </b-col>
      <b-col>
        <b-button variant="danger">Delete</b-button>
      </b-col>
    </b-form-row>
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
                params: {
                  testCollectionInput: testCollection,
                  id: testCollection.id,
                  testGroupId: testGroup.id
                }
              }"
              >Edit</router-link
            >
          </b-col>
          <b-col lg="1">
            <b-button>Execute</b-button>
            <b-button
              variant="danger"
              @click="deleteTestCollection(testCollection.id)"
              >Delete</b-button
            >
          </b-col>
        </b-row>
      </b-list-group-item>
    </b-list-group>
  </div>
</template>

<script>
import { data } from "@/shared/datastore.js";
//const uuid = require("uuid");

export default {
  name: "Tests",
  props: {
    id: {
      type: String,
      default: ""
    },
    testGroupInput: {
      type: Object,
      default: () => {}
    }
  },
  data() {
    return {
      testGroup: { ...this.testGroupInput }
    };
  },
  methods: {
    showDetail(item) {
      if (item.showDetail == true) {
        item.showDetail = false;
      } else {
        item.showDetail = true;
      }
    },
    async deleteTestCollection(id) {
      await data.deleteTestCollection(id);

      let index = this.testGroup.testCollections.findIndex(h => h.id == id);
      this.testGroup.testCollections.splice(index, 1);
    }
  }
};
</script>

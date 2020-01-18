<template>
  <div>
    <b-form>
      <b-form-group label="Name">
        <b-form-input v-model="testCollection.name" />
      </b-form-group>
      <b-form-group label="Description">
        <b-form-input v-model="testCollection.description" />
      </b-form-group>
      <h3>Tests</h3>
      <b-list-group>
        <b-list-group-item v-for="test in testCollection.tests" :key="test.id">
          <b-form>
            <b-form-group label="Name">
              <b-form-input v-model="test.name" />
            </b-form-group>
            <b-form-group label="Method">
              <b-form-select
                v-model="test.method"
                :options="httpMethods"
                size="sm"
                class="mt-3"
              ></b-form-select>
            </b-form-group>
            <b-form-group label="Url">
              <b-form-input v-model="test.url" />
            </b-form-group>
            <b-form-group label="Body">
              <b-form-textarea v-model="test.body" />
            </b-form-group>
            <b-form-group label="Headers">
              <b-list-group>
                <b-list-group-item
                  v-for="header in test.headers"
                  :key="header.id"
                >
                  <b-form-group label="Name">
                    <b-form-input v-model="header.name" />
                  </b-form-group>
                  <b-form-group label="Value">
                    <b-form-input v-model="header.value" />
                  </b-form-group>
                  <b-row>
                    <b-col cols="11"> </b-col>
                    <b-col>
                      <b-button
                        variant="danger"
                        @click="deleteheaderFromTest(test, header)"
                        >Delete</b-button
                      >
                    </b-col>
                  </b-row>
                </b-list-group-item>
              </b-list-group>
              <b-button @click="addHeaderToTest(test)">Add Header</b-button>
            </b-form-group>
            <b-form-group label="Response Code">
              <b-form-input v-model="test.responseCode" type="number" />
            </b-form-group>
            <b-form-group label="Response Body">
              <b-form-textarea v-model="test.responseBody" />
            </b-form-group>
          </b-form>
        </b-list-group-item>
      </b-list-group>
    </b-form>
    <b-row>
      <b-col lg="10" />
      <b-col lg="1">
        <b-button variant="success" @click="saveTestCollection">Save</b-button>
      </b-col>
      <b-col lg="1">
        <router-link
          tag="button"
          class="btn btn-danger"
          :to="{
            name: 'tests',
            params: {
              testGroupId: testGroupId
            }
          }"
        >
          Cancel
        </router-link>
      </b-col>
    </b-row>
  </div>
</template>

<script>
//import { dataStore } from "@/shared/datastoretemp.js";
import { data } from "@/shared/datastore.js";
const uuid = require("uuid");

export default {
  name: "TestDetail",
  props: {
    id: {
      type: String,
      default: "nothing"
    },
    testCollectionInput: {
      type: Object,
      default: () => {}
    },
    testGroupId: {
      type: String,
      default: ""
    }
  },
  data() {
    return {
      testCollection: { ...this.testCollectionInput },
      httpMethods: [
        { value: "GET", text: "GET" },
        { value: "POST", text: "POST" },
        { value: "DELETE", text: "DELETE" },
        { value: "PUT", text: "PUT" },
        { value: "PATCH", text: "PATCH" }
      ]
    };
  },
  methods: {
    async saveTestCollection() {
      this.testCollection.tests.forEach(function(item) {
        item.responseCode = parseInt(item.responseCode);
      });

      await data.updateTestCollection(this.testCollection);
    },
    addHeaderToTest(test) {
      if (test.headers == null) {
        test.headers = [];
      }

      test.headers.push({
        id: uuid(),
        name: "",
        value: ""
      });
    },
    deleteheaderFromTest(test, header) {
      let index = test.headers.findIndex(h => h.id == header.id);
      test.headers.splice(index, 1);
    }
  }
};
</script>

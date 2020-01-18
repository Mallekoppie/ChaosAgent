<template>
  <div>
    <h1>
      Test Groups
    </h1>
    <b-list-group>
      <b-list-group-item v-for="testGroup in testGroups" :key="testGroup.id">
        <b-form>
          <b-form-group label="Name">
            {{ testGroup.name }}
          </b-form-group>
          <b-form-group label="Description">
            {{ testGroup.description }}
          </b-form-group>
        </b-form>
        <b-form-row>
          <b-col cols="11">
            <router-link
              tag="button"
              class="btn"
              :to="{
                name: 'tests',
                params: {
                  testGroupInput: testGroup
                }
              }"
              >View Test Collections</router-link
            >
          </b-col>
          <b-col>
            <b-button variant="danger" @click="deleteTestGroup(testGroup.id)"
              >Delete</b-button
            >
          </b-col>
        </b-form-row>
      </b-list-group-item>
    </b-list-group>
    <b-row>
      <b-col cols="11"> </b-col>
      <b-col cols="1">
        <b-button variant="success" v-b-modal.modal-prevent-closing
          >Add</b-button
        >
      </b-col>
    </b-row>
    <b-modal
      id="modal-prevent-closing"
      ref="modal"
      title="Add Test Group"
      @show="resetModal"
      @hidden="resetModal"
      @ok="handleOk"
    >
      <form ref="form" @submit.stop.prevent="handleSubmit">
        <b-form-group
          :state="nameState"
          label="Name"
          label-for="name-input"
          invalid-feedback="Name is required"
        >
          <b-form-input
            id="name-input"
            v-model="newName"
            :state="nameState"
            required
          ></b-form-input>
        </b-form-group>
        <b-form-group
          :state="nameState"
          label="Description"
          label-for="name-input"
          invalid-feedback="Name is required"
        >
          <b-form-input
            id="description-input"
            v-model="newDescription"
            :state="nameState"
            required
          ></b-form-input>
        </b-form-group>
      </form>
    </b-modal>
  </div>
</template>

<script>
import { dataStore } from "@/shared/datastoretemp.js";
import { data } from "@/shared/datastore.js";
const uuid = require("uuid");

export default {
  name: "Tests",
  data() {
    return {
      testGroups: [],
      nameState: null,
      newName: "",
      newDescription: ""
    };
  },
  async created() {
    this.testGroups = await dataStore.getTestGroups();
    await this.loadTestGroups();
  },
  methods: {
    async loadTestGroups() {
      this.testGroups = [];
      this.testGroups = await data.getAllTestGroups();
    },
    addTestGroup() {
      this.testGroups.push({
        id: uuid(),
        name: "",
        description: "",
        showDetail: false,
        testCollections: []
      });
    },
    async deleteTestGroup(id) {
      await data.deleteTestGroup(id);

      let index = this.testGroups.findIndex(h => h.id == id);
      this.testGroups.splice(index, 1);
    },
    async handleOk(bvModalEvt) {
      bvModalEvt.preventDefault();

      if (this.newName.length > 0 && this.newDescription.length > 0) {
        let newTestGroup = {
          id: uuid(),
          name: this.newName,
          description: this.newDescription,
          showDetail: false,
          testCollections: []
        };

        await data.addTestGroup(newTestGroup);

        this.testGroups.push(newTestGroup);

        // Hide the modal manually
        this.$nextTick(() => {
          this.$bvModal.hide("modal-prevent-closing");
        });

        this.newName = "";
        this.newDescription = "";
      }
    },
    resetModal() {
      this.newName = "";
      this.newDescription = "";
      this.nameState = null;
    }
  }
};
</script>

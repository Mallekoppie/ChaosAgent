<template>
  <div>
    <b-list-group>
      <b-list-group-item v-for="agent in agents" :key="agent.agentId">
        <b-form inline>
          <b-form-group label="Host" class="mb-2 mr-sm-2 mb-sm-0">
            <b-form-input v-model="agent.host" class="mb-2 mr-sm-2 mb-sm-0" />
          </b-form-group>

          <b-form-group label="Port" class="mb-2 mr-sm-2 mb-sm-0">
            <b-form-input v-model="agent.port" class="mb-2 mr-sm-2 mb-sm-0" />
          </b-form-group>

          <b-form-group label="Enabled" class="mb-2 mr-sm-2 mb-sm-0">
            <b-form-checkbox v-model="agent.enabled" switch class="mb-2 mr-sm-2 mb-sm-0" />
          </b-form-group>
          <b-form-group label="Status" class="mb-2 mr-sm-2 mb-sm-0">{{agent.status}}</b-form-group>
          <b-button variant="danger">Delete</b-button>
        </b-form>
      </b-list-group-item>
    </b-list-group>

    <b-row>
      <b-col lg="10" />
      <b-col lg="1">
        <b-button variant="success">Save</b-button>
      </b-col>
      <b-col lg="1">
        <b-button variant="danger">Cancel</b-button>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { dataStore } from "@/shared/datastoretemp.js";

export default {
  name: "Agents",
  data() {
    return {
      agents: []
    };
  },
  created() {
    this.agents = dataStore.getAgents();
  }
};
</script>

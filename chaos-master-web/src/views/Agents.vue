<template>
  <div>
    <b-list-group>
      <b-list-group-item v-for="agent in agents" :key="agent.id">
        <b-form>
          <b-form-group label="Host" class="mb-2 mr-sm-2 mb-sm-0">
            <b-form-input v-model="agent.host" class="mb-2 mr-sm-2 mb-sm-0" />
            <b-form-invalid-feedback :state="validateHostname(agent)"
              >Hostname must be provided</b-form-invalid-feedback
            >
          </b-form-group>

          <b-form-group label="Port" class="mb-2 mr-sm-2 mb-sm-0">
            <b-form-input v-model="agent.port" class="mb-2 mr-sm-2 mb-sm-0" />
            <b-form-invalid-feedback :state="validatePort(agent)">
              Port must be between 2000 and 65200 long.
            </b-form-invalid-feedback>
          </b-form-group>

          <b-form-group label="Metrics Port" class="mb-2 mr-sm-2 mb-sm-0">
            <b-form-input
              v-model="agent.metricsPort"
              class="mb-2 mr-sm-2 mb-sm-0"
            />
            <b-form-invalid-feedback :state="validateMetricsPort(agent)">
              Port must be between 2000 and 65200 long.
            </b-form-invalid-feedback>
          </b-form-group>

          <b-form-group label="Enabled" class="mb-2 mr-sm-2 mb-sm-0">
            <b-form-checkbox
              v-model="agent.enabled"
              switch
              class="mb-2 mr-sm-2 mb-sm-0"
            />
          </b-form-group>
          <b-form-group label="Status" class="mb-2 mr-sm-2 mb-sm-0">{{
            agent.status
          }}</b-form-group>
          <b-form-group label="Functions">
            <b-row>
              <b-col lg="3">
                <b-button variant="success" @click="saveAgent(agent)"
                  >Save</b-button
                >
              </b-col>
              <b-col lg="3">
                <b-button variant="danger" @click="deleteAgent(agent)"
                  >Delete</b-button
                >
              </b-col>
              <b-col lg="6">
                <b-button variant="warning">Stop Tests</b-button>
              </b-col>
            </b-row>
          </b-form-group>
        </b-form>
      </b-list-group-item>
    </b-list-group>

    <b-row>
      <b-col lg="11" />
      <b-col lg="1">
        <b-button variant="success" @click="addAgent">Add</b-button>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { data } from "@/shared/agentstore.js";
const uuid = require("uuid");

export default {
  name: "Agents",
  data() {
    return {
      agents: []
    };
  },
  async created() {
    await this.getAllAgents();
  },
  methods: {
    async getAllAgents() {
      this.agents = await data.getAllAgents();
    },
    addAgent() {
      this.agents.push({
        agentId: uuid(),
        host: "",
        port: 0,
        enabled: true,
        status: "none",
        metricsPort: 0
      });
    },
    async saveAgent(agent) {
      agent.port = parseInt(agent.port);
      agent.metricsPort = parseInt(agent.metricsPort);

      await data.updateAgent(agent);
    },
    async deleteAgent(agent) {
      await data.deleteAgent(agent);

      const index = this.agents.findIndex(a => a.id == agent.id);
      this.agents.splice(index, 1);
    },
    validatePort(agent) {
      let isNumber = !isNaN(agent.port);

      if (isNumber == false) {
        return isNumber;
      }

      let value = parseInt(agent.port);

      if (value < 2000 || value > 65200) {
        return false;
      }

      return true;
    },
        validateMetricsPort(agent) {
      let isNumber = !isNaN(agent.metricsPort);

      if (isNumber == false) {
        return isNumber;
      }

      let value = parseInt(agent.metricsPort);

      if (value < 2000 || value > 65200) {
        return false;
      }

      return true;
    },
    validateHostname(agent) {
      if (agent.host.length < 1) {
        return false;
      } else {
        return true;
      }
    }
  }
};
</script>

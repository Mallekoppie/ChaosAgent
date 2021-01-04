<template>
  <div class="row">
    <div class="col">
      <card type="tasks" :header-classes="{'text-right': isRTL}">
        <template slot="header">
          <h6 class="title d-inline">Test Groups</h6>
          <p class="card-category d-inline">These are just used to group various test collections</p>
          <base-dropdown menu-on-right=""
                         tag="div"
                         title-classes="btn btn-link btn-icon"
                         aria-label="Settings menu"
                         :class="{'float-left': isRTL}">
            <i slot="title" class="tim-icons icon-settings-gear-63"></i>
            <a class="dropdown-item" href="#/add-agent">Add new Agent</a>
          </base-dropdown>
        </template>
        <div class="table-full-width table-responsive">
          <base-table :data="agents"
                      thead-classes="text-primary">
            <template slot-scope="{row}">
              <td>
                <p class="title">{{row.host}}</p>
                <p class="text-muted">Port: {{row.port}}</p>
                <p class="text-muted">Metrics Port: {{row.metricsPort}}</p>
                <p class="text-muted">Enabled: {{row.enabled}}</p>
                <p class="text-muted">Status: {{row.status}}</p>
              </td>
              <td class="td-actions text-right">
                <base-button type="link" aria-label="edit button">
                  <i class="tim-icons icon-pencil"></i>
                </base-button>
              </td>
            </template>
          </base-table>
        </div>
      </card>
    </div>
  </div>
</template>
<script>
import { data } from "@/shared/datastore.js";
import {BaseTable} from '@/components'

  export default {
    components: {
      BaseTable
    },
    data() {
      return {
        agents: []
      }
    },
    async created() {
      await this.getAllAgents();
    },
    computed: {
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
    },
    mounted() {
      this.i18n = this.$i18n;
      if (this.enableRTL) {
        this.i18n.locale = 'ar';
        this.$rtl.enableRTL();
      }
      this.initBigChart(0);
    },
    beforeDestroy() {
      if (this.$rtl.isRTL) {
        this.i18n.locale = 'en';
        this.$rtl.disableRTL();
      }
    }
  };
</script>
<style>
</style>

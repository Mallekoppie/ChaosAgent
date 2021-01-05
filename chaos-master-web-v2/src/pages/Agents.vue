<template>
      <card>
        <template slot="header">
          <div class="row">
            <h6 class="title d-inline col-md-1">Agents</h6>
            <p class="card-category d-inline col-md-10">These are the workers executing the tests</p>
            <base-dropdown menu-on-right=""
                           tag="div"
                           title-classes="btn btn-link btn-icon"
                           aria-label="Settings menu" class="col-md-1 align-right">
              <i slot="title" class="tim-icons icon-settings-gear-63"></i>
              <a class="dropdown-item" href="#/add-agent">Add new Agent</a>
            </base-dropdown>
          </div>

        </template>
        <div class="table-full-width table-responsive">
          <base-table :data="agents"
                      thead-classes="text-primary">
            <template slot-scope="{row}">
              <td>
                <p class="title">{{row.host}}</p>
                <p class="text-muted">Port: {{row.port}}</p>
                <p class="text-muted">Metrics Port: {{row.metricsPort}}</p>
                <p class="text-muted">Enabled: <i class="tim-icons icon-check-2" v-if="row.enabled" /> <i class="tim-icons icon-alert-circle-exc" v-if="!row.enabled" /> </p>
                <p class="text-muted">Status: {{row.status}}</p>
              </td>
              <td class="td-actions text-right">
                <base-button type="link" aria-label="edit button" @click="updateAgent(row)">
                  <i class="tim-icons icon-pencil"></i>
                </base-button>
              </td>
            </template>
          </base-table>
        </div>
      </card>
</template>
<script>
import { data } from "@/shared/agentstore.js";
import { BaseTable} from '@/components'
import router from "@/router";


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
      updateAgent(agent) {
        router.push({name:'add-agent', params:{agentInput:agent}})
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

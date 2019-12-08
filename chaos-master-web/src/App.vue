<template>
  <div>
    <MenuBar />
    <main class="container">
      <router-view />
    </main>
  </div>
</template>

<script>
import MenuBar from "@/components/menu-bar.vue";
export default {
  name: "App",
  components: { MenuBar }
};
</script>

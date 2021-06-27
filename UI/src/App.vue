<template>
  <div class="bg-black-alt font-sans tracking-normal mx-auto h-auto">
    <div>
      <DarkNavBar/>
    </div>
    <div class="pt-20">
      <div class="p-8 flex">
        <div >
          <router-view/>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import DarkNavBar from "@/components/DarkNavBar";

export default {
  components: {DarkNavBar}
}
</script>
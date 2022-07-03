<template>
  <div class="app">
    <div id="nav">
      <Nav/>
    </div>
    <router-view/>
  </div>
</template>

<script>
import Nav from "@/components/Nav"

export default {
  components: {
    Nav
  }
}
</script>
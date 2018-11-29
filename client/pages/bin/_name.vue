<template>
  <div>
    <h1>Inspect {{ $route.params.name }}</h1>


    <div class="columns">
      <div class="column">
        <div
          v-for="(item,key) in items"
          :key="key"
          class="card"
        >      <header class="card-header">
          <p class="card-header-title">
            {{ item.method }} {{ item.requesturi }}
          </p>
          <a 
            href="#" 
            class="card-header-icon" 
            aria-label="more options">
            <span class="icon">
              <i 
                class="fas fa-angle-down" 
                aria-hidden="true"/>
            </span>
          </a>
        </header>
          <div class="card-content">
            <div class="content">
              {{ new Date(item.time).toISOString() }}
              <ul>
                <li
                  v-for="(v,k) in item.header"
                  :key="k">{{ k }}:{{ v }}</li>
              </ul>
            </div>
          </div>
    </div></div></div>
  </div>
</template>

<script>
export default {
  async fetch({ store, params }) {
    await store.dispatch('bin/fetch_bin', params.name)
  },
  computed: {
    items() {
      return this.$store.state.bin.items[this.$route.params.name]
    }
  }
}
</script>
<style>
.card {
  margin: 1em;
}
</style>

<template>
  <section class="section">
    <div class="field is-horizontal binurl">
      <span>https://</span>
      <div class="control has-icons-right">
        <input 
          v-model="binname"
          class="input"
          type="text"
        >
      </div>
      <span>.{{ host }}</span>

    </div>
    <div class="field">
      <button 
        class="button is-primary is-right" 
        @click="$store.dispatch('createBin', binname)">CreateBin</button>
    </div>

    <div v-if="created">
      {{ error }}
    </div>
  </section>
</template>

<script>
import Logo from '~/components/Logo.vue'

export default {
  components: {
    Logo
  },
  computed: {
    binname: {
      get() {
        return this.$store.state.binname
      },
      set(value) {
        this.$store.commit('updateBinName', value)
      }
    },
    host() {
      return location.host
    },
    protocol() {
      return location.protocol
    },
    created() {
      return this.$store.state.created
    }
  }
}
</script>

<style>
.container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.title {
  font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}

section {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.binurl {
  display: flex;
  align-items: center;
}
</style>

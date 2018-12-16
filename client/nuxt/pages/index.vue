<template>
  <section class="section">
    <div class="field is-horizontal binurl">
      <span>{{ protocol }}//</span>
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

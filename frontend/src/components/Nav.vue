import { computed } from 'vue'; import { computed } from 'vue';
<template>
  <nav class="navbar navbar-expand-md navbar-dark bg-dark mb-4">
    <div class="container-fluid">
      <router-link to="/" class="navbar-brand">Home</router-link>

      <div class="collapse navbar-collapse" id="navbarCollapse">
        <ul v-if="!auth" class="navbar-nav me-auto mb-2 mb-md-0">
          <li class="nav-item">
            <router-link to="/login" class="nav-link">Login</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/register" class="nav-link">Register</router-link>
          </li>
        </ul>
        <ul v-else class="navbar-nav me-auto mb-2 mb-md-0">
          <li class="nav-item">
            <router-link to="/login" class="nav-link"
              ><a href="#" class="nav-link" @click="logout">Logout</a>
            </router-link>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script lang="ts">
import { computed } from "vue";
import { useStore } from "vuex";

export default {
  name: "NavCOmponent",
  setup() {
    const store = useStore();
    const auth = computed(() => store.state.authenticated);
    const logout = async () => {
      try {
        fetch("http://localhost:8000/api/logout", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
        });
        await store.dispatch("setAuth", false);
      } catch (e) {
        await store.dispatch("setAuth", true);
      }
    };

    return { auth, logout };
  },
};
</script>

<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
  </div>
  <div>
    <button @click="healthCheck()">Health Check</button>
  </div>
  <div>
    <span v-if="data && data.status">{{ data.status }}</span>
  </div>
</template>

<script>
import { ref } from "vue";
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  setup() {
    const data = ref({});
    return {
      data,
    };
  },
  methods: {
    healthCheck() {
      var self = this;
      return fetch('/api/health', {
          method: 'GET',
          headers: new Headers({
              'Accept': 'application/json'
          })
      })
      .then(function(response) {
          if (response.status == 200) {
              response.json().then(json => {
                self.data.status = (json.UP ? "UP" : "DOWN");
              });
          } else {
              self.data.status = "DOWN (" + response.statusText + ")";
          }
          return response;
      })
      .catch(function(err) {
        self.data.status = "DOWN (" + err + ")";
      });
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>

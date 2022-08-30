<template>
<v-app>
    <v-card>
        <v-col class="text-center"  sm="15">
    <v-img aspect-ratio="10" src="../assets/logo.png"/> 
    <div>
    <v-btn variant="outlined" color="Blue">AboutUS</v-btn> 
    </div>
    <div>
    <v-btn  to = '/login' flat color="secondary">Login</v-btn>
    </div>
        </v-col>
    </v-card>
    
    <v-main>
    
  <ImgComparisonSlider class="slider-example-relative-size">
    <!-- eslint-disable -->
    <img
      slot="first"
      style="width: 100%"
      src="https://i.postimg.cc/8CzG2sbb/pp.png"
    />
    <img
      slot="second"
      style="width: 100%"
      src="https://i.postimg.cc/gkWC9nz6/try1.jpg"
    />
    <!-- eslint-enable -->
    </ImgComparisonSlider>
    </v-main>
  </v-app>
</template>

<style>
  .slider-example-relative-size {
    --default-handle-width: clamp(40px, 10vw, 200px);
  }
</style>

<script>
import { ImgComparisonSlider } from '@img-comparison-slider/vue';
import axios from "axios"

  export default {
    name: 'ExampleComponent',
  components: {
    ImgComparisonSlider,
  },
    async mounted() {
      const url = "http://localhost:5000/user"
      await axios.get(url, {
        withCredentials: true
      }).then((response) => {
        if (response.status === 200) {
          this.user = response.data.name
          this.email = response.data.email
        }
      })
    },
    data() {
      return {
        user: "",
        email: ""
      }
    }
  }
</script>
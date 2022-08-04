<template>
  <header>
    <div class="text-center bg-black text-gray-800 py-10 px-6">
      <h1 class="text-5xl font-bold mt-0 mb-6 title">{{ title }}</h1>
      <h3 class="text-2xl mb-8 font-semibold text-white">{{ subTitle }}</h3>
    </div>
  </header>

  <div class="container grid w-screen place-items-center mx-auto my-10 py-5">
    <form class="w-3/4 max-w-sm">
      <div class="flex items-center border-b border-black py-2">
        <input
          class="appearance-none bg-transparent border-none w-full text-black mr-3 py-1 px-2 leading-tight focus:outline-none placeholder-black"
          type="url"
          placeholder="Enter/Paste your URL ..."
          aria-label="url"
          v-model="inputURL"
        />
        <button
          class="flex-shrink-0 bg-black hover:bg-black hover:border-black border-black text-sm border-4 text-white py-1 px-2 rounded transition ease-in-out delay-50 bg-black hover:-translate-y-1 hover:scale-110 hover:bg-black duration-300 ..."
          type="button"
          @click="fetchShortURL"
        >
          GO!
        </button>
      </div>
    </form>
  </div>

  <div
    class="container grid w-screen place-items-center mx-auto my-10 text-1xl"
    v-if="showResponse"
  >
    <div class="container grid w-screen place-items-center mx-auto my-5">
      <qrcode-vue
        :value="response.custom_short"
        :size="qrCodeSize"
      ></qrcode-vue>
    </div>

    <div class="container grid w-screen place-items-center mx-2">
      <span>
        Scan this QR to get your shortened URL. Please note that this is valid
        for
        <span class="underline">{{ response.expiry }} minutes </span>
        only.
      </span>
    </div>
  </div>

  <div
    class="container grid w-3/4 place-items-center mx-auto text-center"
    v-if="showInputAlert"
  >
    <div
      class="flex p-4 mb-4 text-sm rounded-lg dark:bg-black dark:text-white"
      role="alert"
    >
      <svg
        aria-hidden="true"
        class="flex-shrink-0 inline w-5 h-5 mr-3"
        fill="currentColor"
        viewBox="0 0 20 20"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          fill-rule="evenodd"
          d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
          clip-rule="evenodd"
        ></path>
      </svg>
      <span class="sr-only">Info</span>
      <div>
        <span class="font-medium">Info Alert!</span> Please input a valid URL to
        continue.
      </div>
    </div>
  </div>
</template>

<script>
import QrcodeVue from "qrcode.vue";
import axios from "axios";

export default {
  name: "App",
  components: {
    QrcodeVue,
  },
  data() {
    return {
      title: "GoShort",
      subTitle: "Shorten Your URLs In One GO",
      inputURL: "",
      showResponse: false,
      qrCodeSize: 130,
      response: null,
      shortenURLEndpoint: "localhost:3000/shorten",
    };
  },
  computed: {
    showInputAlert() {
      if (this.inputURL.length <= 10) {
        return true;
      }
      return false;
    },
  },
  methods: {
    fetchShortURL() {
      axios
        .post(shortenURLEndpoint, { url: this.inputURL })
        .then((res) => {
          this.response = res.data;
          var audio = new Audio(require("@/assets/ping.mp3"));
          audio.play();

          if (!this.showInputAlert) {
            this.showResponse = true;
          }
        })
        .catch((err) => {
          this.showErrorAlert = true;
          console.error(err);
        });
    },
  },
  created() {
    document.title = "GoShort";
  },
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Righteous&display=swap");

* {
  font-family: "Righteous", cursive;
}
.title {
  background: -webkit-linear-gradient(rgb(249, 253, 253), rgb(172, 165, 165));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.underline {
  text-decoration: underline;
}
</style>

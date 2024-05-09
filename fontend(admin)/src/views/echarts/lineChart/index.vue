<template>
  <div>
    <div class="tabs">
      <button @click="showCard('card1')" :class="[tabActive == 1 ? 'tabactive' : '']">Week</button>
      <button @click="showCard('card2')" :class="[tabActive == 2 ? 'tabactive' : '']">Month</button>
      <button @click="showCard('card3')" :class="[tabActive == 3 ? 'tabactive' : '']">Year</button>
    </div>

    <div id="card1" class="card content-box active">
      <ECharts :option="option1" />
    </div>

    <div id="card2" class="card content-box mynone">
      <ECharts :option="option2" />
    </div>

    <div id="card3" class="card content-box mynone">
      <ECharts :option="option3" />
    </div>
  </div>
</template>

<script setup lang="ts" name="lineChart">
import { ECOption } from "@/components/ECharts/config";
import ECharts from "../index.vue";
import { ref } from "vue";
let tabActive = ref(1);
function showCard(cardId) {
  tabActive.value = cardId[4];
  console.log(tabActive);
  const cards = document.querySelectorAll(".card");
  cards.forEach(card => {
    if (card.id === cardId) {
      card.classList.add("myactive");
      card.classList.remove("mynone");
    } else {
      card.classList.remove("myactive");
      card.classList.add("mynone");
    }
  });
}
const option1: ECOption = {
  title: {
    text: "Future revenue for one week",
    textStyle: {
      color: "#a1a1a1"
    }
  },
  tooltip: {
    trigger: "axis",
    axisPointer: {
      type: "cross",
      label: {
        backgroundColor: "#6a7985"
      }
    }
  },
  legend: {
    data: ["Revenue"],
    textStyle: {
      color: "#a1a1a1"
    }
  },
  toolbox: {
    feature: {
      saveAsImage: {}
    }
  },
  grid: {
    left: "3%",
    right: "4%",
    bottom: "3%",
    containLabel: true
  },
  xAxis: [
    {
      type: "category",
      boundaryGap: false,
      data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
      axisLabel: {
        color: "#a1a1a1"
      }
    }
  ],
  yAxis: [
    {
      type: "value",
      axisLabel: {
        color: "#a1a1a1"
      }
    }
  ],
  series: [
    {
      name: "Revenue",
      type: "line",
      stack: "Total",
      areaStyle: {},
      emphasis: {
        focus: "series"
      },
      data: [120, 132, 101, 134, 90, 230, 210]
    }
  ]
};
const option2: ECOption = {
  title: {
    text: "Future revenue for one month",
    textStyle: {
      color: "#a1a1a1"
    }
  },
  tooltip: {
    trigger: "axis",
    axisPointer: {
      type: "cross",
      label: {
        backgroundColor: "#6a7985"
      }
    }
  },
  legend: {
    data: ["Revenue"],
    textStyle: {
      color: "#a1a1a1"
    }
  },
  toolbox: {
    feature: {
      saveAsImage: {}
    }
  },
  grid: {
    left: "3%",
    right: "4%",
    bottom: "3%",
    containLabel: true
  },
  xAxis: [
    {
      type: "category",
      boundaryGap: false,
      data: ["week 1", "week 2", "week 3", "week 4"],
      axisLabel: {
        color: "#a1a1a1"
      }
    }
  ],
  yAxis: [
    {
      type: "value",
      axisLabel: {
        color: "#a1a1a1"
      }
    }
  ],
  series: [
    {
      name: "Revenue",
      type: "line",
      stack: "Total",
      color: "rgba(237, 176, 63, 0.75)",
      areaStyle: {},
      emphasis: {
        focus: "series"
      },
      data: [1085, 2898, 2101, 2453]
    }
  ]
};
const option3: ECOption = {
  title: {
    text: "Future revenue for one year",
    textStyle: {
      color: "#a1a1a1"
    }
  },
  tooltip: {
    trigger: "axis",
    axisPointer: {
      type: "cross",
      label: {
        backgroundColor: "#6a7985"
      }
    }
  },
  legend: {
    data: ["Revenue"],
    textStyle: {
      color: "#a1a1a1"
    }
  },
  toolbox: {
    feature: {
      saveAsImage: {}
    }
  },
  grid: {
    left: "3%",
    right: "4%",
    bottom: "3%",
    containLabel: true
  },
  xAxis: [
    {
      type: "category",
      boundaryGap: false,
      data: ["June", "July", "August", "September", "October", "November", "December"],
      axisLabel: {
        color: "#a1a1a1"
      }
    }
  ],
  yAxis: [
    {
      type: "value",
      axisLabel: {
        color: "#a1a1a1"
      }
    }
  ],
  series: [
    {
      name: "Revenue",
      type: "line",
      stack: "Total",
      areaStyle: {},
      color: "rgba(225, 64, 64, 0.65)",
      emphasis: {
        focus: "series"
      },
      data: [12453, 12564, 21765, 22423, 21675, 19563, 23534]
    }
  ]
};
</script>

<style scoped lang="scss">
@import "./index.scss";
.myactive {
  display: block;
}
.mynone {
  display: none;
}
#card1,
#card2,
#card3 {
  height: 600px;
}
.tabs button {
  margin-right: 10px;
  width: 80px;
  background-color: white;
  border-radius: 5px;
  font-size: 20px;
  height: 30px;
}
.tabactive {
  background-color: rgb(30, 59, 134) !important;
  color: white;
}
</style>

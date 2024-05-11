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
// import { ECOption } from "@/components/ECharts/config";
import ECharts from "../index.vue";
// import { ref } from "vue";
import { ref, onMounted } from "vue";
import axios from "axios";

// 为每个图表卡片准备 option，使用 ref 创建响应式引用
const option1 = ref({});
const option2 = ref({});
const option3 = ref({});

// 初始 ECharts 配置项模板
const baseOption = {
  title: {
    text: "",
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
  xAxis: {
    type: "category",
    boundaryGap: false,
    data: [],
    axisLabel: {
      color: "#a1a1a1"
    }
  },
  yAxis: {
    type: "value",
    axisLabel: {
      color: "#a1a1a1"
    }
  },
  series: [
    {
      name: "Revenue",
      type: "line",
      stack: "Total",
      areaStyle: {},
      emphasis: {
        focus: "series"
      },
      data: [] // 后端传递
    }
  ]
};

// 设置初始的 xAxis 数据
const xAxisDataWeekly = ["Current Week", "The Second Week", "The Third Week", "The Fourth Week"];
const xAxisDataMonthly = ["Current Month", "The Second Month", "The Third Month", "The Fourth Month"];
const xAxisDataYearly = ["Current Year", "The Second Year", "The Third Year"];

// 初始化卡片的 ECharts 配置项
onMounted(() => {
  option1.value = { ...baseOption, title: { text: "Future revenue for one week" }, xAxis: { data: xAxisDataWeekly } };
  option2.value = { ...baseOption, title: { text: "Future revenue for one month" }, xAxis: { data: xAxisDataMonthly } };
  option3.value = { ...baseOption, title: { text: "Future revenue for one year" }, xAxis: { data: xAxisDataYearly } };
  // 获取后端数据并更新 series
  fetchData("http://120.46.81.37:8080/admin/week", option1);
  fetchData("http://120.46.81.37:8080/admin/month", option2);
  fetchData("http://120.46.81.37:8080/admin/year", option3);
});

// 获取数据函数
const fetchData = async (url, optionRef) => {
  try {
    const response = await axios.get(url);
    optionRef.value.series[0].data = response.data; // 假设返回的数据是一个数组
  } catch (error) {
    console.error("Error fetching data:", error);
  }
};

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
// const option1: ECOption = {
//   title: {
//     text: "Future revenue for one week",
//     textStyle: {
//       color: "#a1a1a1"
//     }
//   },
//   tooltip: {
//     trigger: "axis",
//     axisPointer: {
//       type: "cross",
//       label: {
//         backgroundColor: "#6a7985"
//       }
//     }
//   },
//   legend: {
//     data: ["Revenue"],
//     textStyle: {
//       color: "#a1a1a1"
//     }
//   },
//   toolbox: {
//     feature: {
//       saveAsImage: {}
//     }
//   },
//   grid: {
//     left: "3%",
//     right: "4%",
//     bottom: "3%",
//     containLabel: true
//   },
//   xAxis: [
//     {
//       type: "category",
//       boundaryGap: false,
//       data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
//       axisLabel: {
//         color: "#a1a1a1"
//       }
//     }
//   ],
//   yAxis: [
//     {
//       type: "value",
//       axisLabel: {
//         color: "#a1a1a1"
//       }
//     }
//   ],
//   series: [
//     {
//       name: "Revenue",
//       type: "line",
//       stack: "Total",
//       areaStyle: {},
//       emphasis: {
//         focus: "series"
//       },
//       data: [120, 132, 101, 134, 90, 230, 210]
//     }
//   ]
// };
// const option2: ECOption = {
//   title: {
//     text: "Future revenue for one month",
//     textStyle: {
//       color: "#a1a1a1"
//     }
//   },
//   tooltip: {
//     trigger: "axis",
//     axisPointer: {
//       type: "cross",
//       label: {
//         backgroundColor: "#6a7985"
//       }
//     }
//   },
//   legend: {
//     data: ["Revenue"],
//     textStyle: {
//       color: "#a1a1a1"
//     }
//   },
//   toolbox: {
//     feature: {
//       saveAsImage: {}
//     }
//   },
//   grid: {
//     left: "3%",
//     right: "4%",
//     bottom: "3%",
//     containLabel: true
//   },
//   xAxis: [
//     {
//       type: "category",
//       boundaryGap: false,
//       data: ["week 1", "week 2", "week 3", "week 4"],
//       axisLabel: {
//         color: "#a1a1a1"
//       }
//     }
//   ],
//   yAxis: [
//     {
//       type: "value",
//       axisLabel: {
//         color: "#a1a1a1"
//       }
//     }
//   ],
//   series: [
//     {
//       name: "Revenue",
//       type: "line",
//       stack: "Total",
//       color: "rgba(237, 176, 63, 0.75)",
//       areaStyle: {},
//       emphasis: {
//         focus: "series"
//       },
//       data: [1085, 2898, 2101, 2453]
//     }
//   ]
// };
// const option3: ECOption = {
//   title: {
//     text: "Future revenue for one year",
//     textStyle: {
//       color: "#a1a1a1"
//     }
//   },
//   tooltip: {
//     trigger: "axis",
//     axisPointer: {
//       type: "cross",
//       label: {
//         backgroundColor: "#6a7985"
//       }
//     }
//   },
//   legend: {
//     data: ["Revenue"],
//     textStyle: {
//       color: "#a1a1a1"
//     }
//   },
//   toolbox: {
//     feature: {
//       saveAsImage: {}
//     }
//   },
//   grid: {
//     left: "3%",
//     right: "4%",
//     bottom: "3%",
//     containLabel: true
//   },
//   xAxis: [
//     {
//       type: "category",
//       boundaryGap: false,
//       data: ["June", "July", "August", "September", "October", "November", "December"],
//       axisLabel: {
//         color: "#a1a1a1"
//       }
//     }
//   ],
//   yAxis: [
//     {
//       type: "value",
//       axisLabel: {
//         color: "#a1a1a1"
//       }
//     }
//   ],
//   series: [
//     {
//       name: "Revenue",
//       type: "line",
//       stack: "Total",
//       areaStyle: {},
//       color: "rgba(225, 64, 64, 0.65)",
//       emphasis: {
//         focus: "series"
//       },
//       data: [12453, 12564, 21765, 22423, 21675, 19563, 23534]
//     }
//   ]
// };
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
  width: 80px;
  height: 30px;
  margin-right: 10px;
  font-size: 20px;
  background-color: white;
  border-radius: 5px;
}
.tabactive {
  color: white;
  background-color: rgb(30 59 134) !important;
}
</style>

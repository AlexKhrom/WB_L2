<template>
  <v-app style="padding: 20px;">
    <template v-if="isNewEvent">
      <div class="picker" @click="isNewEvent=false">
        <div class="picker__content" @click.stop>
          <slot>
            <v-date-picker
                :min="(new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)"
                v-model="picker"
                header-color="black primary"
            ></v-date-picker>

            <v-text-field
                v-model="eventTitle"
                style="margin-top: 30px"
                solo
                label="title"
            ></v-text-field>
            <v-btn
                style="width: 100%"
                depressed
                color="primary"
                @click="creatEvent"
            >
              creat
            </v-btn>

          </slot>
        </div>
      </div>
    </template>


    <div style="width: 270px;margin: auto;margin-top: 30px">
      <slot>
        <v-btn
            style="width: 100%"
            v-if="!isNewEvent"
            @click="isNewEvent=true"
            color="primary"
        >
          new event
        </v-btn>
        <v-btn
            style="width: 100%;margin-top: 30px"
            @click="getEventsDay()"
            color="primary"
        >
          events for day
        </v-btn>
        <v-btn
            style="width: 100%;margin-top: 30px"
            @click="getEventsWeek()"
            color="primary"
        >
          events for week
        </v-btn>
        <v-btn
            style="width: 100%;margin-top: 30px"
            @click="getEventsMonth()"
            color="primary"
        >
          events for month
        </v-btn>
      </slot>
    </div>

    <div style="width: 350px;margin: auto;margin-top: 50px">
      <h1>{{duration}}</h1>
      <!--      <div v-for="(event) in this.events"-->
      <!--           :key="event.id">-->
      <!--        <h1>{{event.title}}</h1><h2>{{(new Date(event.date - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)}}</h2>-->
      <!--      </div>-->

      <div v-for="(event) in this.events"
           :key="event.id"
      >
        <!--        <div v-if="isNewDay(index)" style="margin-top: 20px;color: #555555">{{ new Date(event.dt) | moment }}</div>-->
        <div
            style="display: flex;justify-content: space-between;margin-top: 20px;padding:5px;">
          <div>
            <div :style="event.deleted?'opacity: 0.5':''"
                 style="font-size: 18px;text-transform: capitalize;display: flex;align-items: center;">{{ event.title }}
            </div>
            <!--            <span style="color: #555;font-size: 11px;display: block">{{-->
            <!--                (new Date(event.date - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)-->
            <!--              }}</span>-->
          </div>
          <div>
             <span style="color: #555;font-size: 18px;">{{
                 (new Date(event.date - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)
               }}</span>
            <v-btn @click="changeEvent(event)" icon>✏️</v-btn>
            <v-btn @click="deleteEvent(event)" icon>🗑</v-btn>
          </div>
        </div>
        <!--        <div :style="{border: '2px solid '+borderColorForEvent(index),width: calcTimePercent(index) +'%'}"-->
        <!--        ></div>-->
      </div>
    </div>

    <template>
      <v-row justify="center">
        <v-dialog
            v-model="isDeleteEvent"
            persistent
            max-width="290"
        >
          <v-card>
            <v-card-title class="text-h5">
              Delete {{ deleteEventObj.title }} ?
            </v-card-title>
            <!--            <v-card-text>Let Google help apps determine location. This means sending anonymous location data to Google,-->
            <!--              even when no apps are running.-->
            <!--            </v-card-text>-->
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                  color="darken-1"
                  text
                  @click="isDeleteEvent = false;"
              >
                No
              </v-btn>
              <v-btn
                  color="green darken-1"
                  text
                  @click="deleteEventWithId()"
              >
                Yes
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-row>
    </template>

    <template v-if="isChangeEvent">
      <div class="picker" @click="isChangeEvent=false">
        <div class="picker__content" @click.stop>
          <slot>
            <v-date-picker
                v-model="changePicker"
                header-color="black primary"
            ></v-date-picker>

            <v-text-field
                v-model="changeEventObj.title"
                style="margin-top: 30px"
                solo
                label="title"
            ></v-text-field>
            <v-btn
                style="width: 100%"
                depressed
                color="primary"
                @click="saveChangeEvent()"
            >
              save
            </v-btn>

          </slot>
        </div>
      </div>
    </template>
  </v-app>
</template>

<script>


export default {
  name: 'App',
  metaInfo: {
    title: 'tasks'
  },
  data() {
    return {
      duration: "",

      picker: (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10),
      isNewEvent: false,
      eventTitle: "",
      newEvent: {
        userId: 0,
        title: "",
        date: -1,
      },
      isChangeEvent: false,
      changeEventId: -1,
      changeEventObj: {},
      changePicker: (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10),

      isDeleteEvent: false,
      deleteEventId: -1,
      deleteEventObj: {},

      events: [],

      lastCommand: function () {
      },
    }
  },
  methods: {
    async creatEvent() {
      this.isNewEvent = !this.isNewEvent
      this.newEvent.userId = 0
      this.newEvent.title = this.eventTitle
      this.newEvent.date = (new Date(this.picker)).getTime()

      console.log("newTask = ", this.newEvent)
      console.log(this.newEvent.date)
      console.log(this.picker)

      this.eventTitle = ""
      this.picker = (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)

      this.makeReq("/api/createEvent", "POST", JSON.stringify(this.newEvent))

      this.newEvent = {
        userId: 0,
        title: "",
        date: -1,
      }
      setTimeout(this.lastCommand, 100)
    },

    async getEventsDay() {
      let response = await this.makeReq("/api/events_for_day?user_id="+this.newEvent.userId+"&date=" + (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10), "GET")
      let respBody = await response.json()
      console.log("resp body = ", respBody)

      this.events = respBody
      this.events.sort(this.compare)
      this.lastCommand = this.getEventsDay
      this.duration = "Events for day"
    },

    async getEventsWeek() {
      let response = await this.makeReq("/api/events_for_week?user_id="+this.newEvent.userId+"&date=" + (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10), "GET")
      let respBody = await response.json()
      console.log("resp body = ", respBody)
      this.events = respBody
      this.events.sort(this.compare)
      this.lastCommand = this.getEventsWeek
      this.duration = "Events for week"
    },
    async getEventsMonth() {
      let response = await this.makeReq("/api/events_for_month?user_id="+this.newEvent.userId+"&date=" + (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10), "GET")
      let respBody = await response.json()
      console.log("resp body = ", respBody)
      this.events = respBody
      this.events.sort(this.compare)
      this.lastCommand = this.getEventsMonth
      this.duration = "Events for month"
    },

    changeEvent(event) {
      this.changeEventObj = event
      this.isChangeEvent = true
      this.changePicker = (new Date(event.date - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)
    },

    async saveChangeEvent() {
      this.changeEventObj.date = (new Date(this.changePicker)).getTime()
      this.makeReq("/api/updateEvent", "UPDATE", JSON.stringify(this.changeEventObj))
      this.isChangeEvent = false
      setTimeout(this.lastCommand, 100)
      this.changeEventObj = {}

    },

    deleteEvent(event) {
      this.deleteEventObj = event
      this.isDeleteEvent = true
      this.lastCommand()
    },

    async deleteEventWithId() {
      this.makeReq("/api/deleteEvent", "DELETE", JSON.stringify(this.deleteEventObj))
      setTimeout(this.lastCommand, 100)
      this.deleteEventObj = {}
      this.isDeleteEvent = false;
    },

    async makeReq(url, method, body) {
      let response
      if (body === undefined) {
        response = await fetch(url, {
          method: method,
        });
      } else {
        response = await fetch(url, {
          method: method,
          body: body
        });
      }

      if (response.ok) { // если HTTP-статус в диапазоне 200-299
        // получаем тело ответа (см. про этот метод ниже)
        console.log("okkkk")
        return response
      } else if (response.status > 299 && response.status < 500) {
        console.log("err")
        return 'err'
      } else if (response.status > 500) {
        console.log("some wrong on backend")
        return 'backend err'
      }
    },
    compare(a, b) {
      if (a.date < b.date) {
        return -1;
      }
      if (a.date > b.date) {
        return 1;
      }
      return 0;
    }

  },
  beforeMount() {
    this.getEventsDay()
  },
}
</script>

<style scoped>

.picker {
  color: white;
  top: 0;
  bottom: 0;
  right: 0;
  left: 0;
  background-color: rgba(0, 0, 0, 0.5);
  position: fixed;
  z-index: 100;
  overflow: auto;
  display: flex;
  padding-top: 80px;
}

.picker__content {
  /*overflow-y: scroll;*/
  display: block;
  width: 290px;
  margin: auto;


  min-height: 100%;
  justify-content: space-between;
  z-index: 200;
}

</style>

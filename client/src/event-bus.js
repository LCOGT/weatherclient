/*

Apparently using an Event Bus is the only way to share data between independent/non-related components.
See: https://alligator.io/vuejs/global-event-bus/
 */

import Vue from 'vue';

export const EventBus= new Vue();

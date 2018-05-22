import Vue from 'vue';

Vue.filter('statusMsgToLetter', msg => (msg === 'true' | msg === 'Unknown') ? 'Y': 'N');

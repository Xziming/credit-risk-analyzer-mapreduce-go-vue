import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
history: createWebHistory(import.meta.env.BASE_URL),
scrollBehavior(to, from, savedPosition) {
return savedPosition || { left: 0, top: 0 }
},
routes: [
{
path: '/',
name: 'Ecommerce',
component: () => import('../views/Ecommerce.vue'),
meta: {
title: 'eCommerce Dashboard',
},
},
{
path: '/calendar',
name: 'Calendar',
component: () => import('../views/Others/Calendar.vue'),
meta: {
title: 'Calendar',
},
},
{
path: '/profile',
name: 'Profile',
component: () => import('../views/Others/UserProfile.vue'),
meta: {
title: 'Profile',
},
},
{
path: '/form-elements',
name: 'Form Elements',
component: () => import('../views/Forms/FormElements.vue'),
meta: {
title: 'Form Elements',
},
},
{
path: '/basic-tables',
name: 'Basic Tables',
component: () => import('../views/Tables/BasicTables.vue'),
meta: {
title: 'Basic Tables',
},
},
{
path: '/line-chart',
name: 'Line Chart',
component: () => import('../views/Chart/LineChart/LineChart.vue'),
},
{
path: '/bar-chart',
name: 'Bar Chart',
component: () => import('../views/Chart/BarChart/BarChart.vue'),
},
{
path: '/age-chart',
name: 'Age Chart',
component: () => import('../views/Chart/BasicChart/AgeChart.vue'),
},
{
path: '/age-d-chart',
name: 'Age Defaults Chart',
component: () => import('../views/Chart/DefaultChart/AgeDefaultsChart.vue'),
},
{
path: '/gender-chart',
name: 'Gender Chart',
component: () => import('../views/Chart/BasicChart/GenderChart.vue'),
},
{
path: '/gender-d-chart',
name: 'Gender Defaults Chart',
component: () => import('../views/Chart/DefaultChart/GenderDefaultsChart.vue'),
},
{
path: '/education-chart',
name: 'Education Chart',
component: () => import('../views/Chart/BasicChart/EducationChart.vue'),
},
{
path: '/education-d-chart',
name: 'Education Defaults Chart',
component: () => import('../views/Chart/DefaultChart/EducationDefaultsChart.vue'),
},
{
path: '/marriage-chart',
name: 'Marriage Chart',
component: () => import('../views/Chart/BasicChart/MarriageChart.vue'),
},
{
path: '/marriage-d-chart',
name: 'Marriage Defaults Chart',
component: () => import('../views/Chart/DefaultChart/MarriageDefaultsChart.vue'),
},
{
path: '/limit-chart',
name: 'Limit Bal Chart',
component: () => import('../views/Chart/BasicChart/LimitBalChart.vue'),
},
{
path: '/limit-d-chart',
name: 'Limit Bal Defaults Chart',
component: () => import('../views/Chart/DefaultChart/LimitBalDefaultsChart.vue'),
},
{
path: '/alerts',
name: 'Alerts',
component: () => import('../views/UiElements/Alerts.vue'),
meta: {
title: 'Alerts',
},
},
{
path: '/avatars',
name: 'Avatars',
component: () => import('../views/UiElements/Avatars.vue'),
meta: {
title: 'Avatars',
},
},
{
path: '/badge',
name: 'Badge',
component: () => import('../views/UiElements/Badges.vue'),
meta: {
title: 'Badge',
},
},

{
path: '/buttons',
name: 'Buttons',
component: () => import('../views/UiElements/Buttons.vue'),
meta: {
title: 'Buttons',
},
},

{
path: '/images',
name: 'Images',
component: () => import('../views/UiElements/Images.vue'),
meta: {
title: 'Images',
},
},

{
path: '/alerts',
name: 'Alerts',
component: () => import('../views/UiElements/Alerts.vue'),
meta: {
title: 'Alerts',
},
},
{
path: '/avatars',
name: 'Avatars',
component: () => import('../views/UiElements/Avatars.vue'),
meta: {
title: 'Avatars',
},
},
{
path: '/badge',
name: 'Badge',
component: () => import('../views/UiElements/Badges.vue'),
meta: {
title: 'Badge',
},
},

{
path: '/buttons',
name: 'Buttons',
component: () => import('../views/UiElements/Buttons.vue'),
meta: {
title: 'Buttons',
},
},

{
path: '/images',
name: 'Images',
component: () => import('../views/UiElements/Images.vue'),
meta: {
title: 'Images',
},
},

{
path: '/alerts',
name: 'Alerts',
component: () => import('../views/UiElements/Alerts.vue'),
meta: {
title: 'Alerts',
},
},

{
path: '/avatars',
name: 'Avatars',
component: () => import('../views/UiElements/Avatars.vue'),
meta: {
title: 'Avatars',
},
},

{
path: '/badge',
name: 'Badge',
component: () => import('../views/UiElements/Badges.vue'),
meta: {
title: 'Badge',
},
},

{
path: '/buttons',
name: 'Buttons',
component: () => import('../views/UiElements/Buttons.vue'),
meta: {
title: 'Buttons',
},
},

{
path: '/images',
name: 'Images',
component: () => import('../views/UiElements/Images.vue'),
meta: {
title: 'Images',
},
},

{
path: '/alerts',
name: 'Alerts',
component: () => import('../views/UiElements/Alerts.vue'),
meta: {
title: 'Alerts',
},
},

{
path: '/avatars',
name: 'Avatars',
component: () => import('../views/UiElements/Avatars.vue'),
meta: {
title: 'Avatars',
},
},

{
path: '/badge',
name: 'Badge',
component: () => import('../views/UiElements/Badges.vue'),
meta: {
title: 'Badge',
},
},

{
path: '/buttons',
name: 'Buttons',
component: () => import('../views/UiElements/Buttons.vue'),
meta: {
title: 'Buttons',
},
},

{
path: '/images',
name: 'Images',
component: () => import('../views/UiElements/Images.vue'),
meta: {
title: 'Images',
},
},

{
path: '/alerts',
name: 'Alerts',
component: () => import('../views/UiElements/Alerts.vue'),
meta: {
title: 'Alerts',
},
},

{
path: '/avatars',
name: 'Avatars',
component: () => import('../views/UiElements/Avatars.vue'),
meta: {
title: 'Avatars',
},
},

{
path: '/badge',
name: 'Badge',
component: () => import('../views/UiElements/Badges.vue'),
meta: {
title: 'Badge',
},
},

{
path: '/buttons',
name: 'Buttons',
component: () => import('../views/UiElements/Buttons.vue'),
meta: {
title: 'Buttons',
},
},

{
path: '/images',
name: 'Images',
component: () => import('../views/UiElements/Images.vue'),
meta: {
title: 'Images',
},
},

{
path: '/videos',
name: 'Videos',
component: () => import('../views/UiElements/Videos.vue'),
meta: {
title: 'Videos',
},
},

{
path: '/blank',
name: 'Blank',
component: () => import('../views/Pages/BlankPage.vue'),
meta: {
title: 'Blank',
},
},

{
path: '/error-404',
name: '404 Error',
component: () => import('../views/Errors/FourZeroFour.vue'),
meta: {
title: '404 Error',
},
},

{
path: '/signin',
name: 'Signin',
component: () => import('../views/Auth/Signin.vue'),
meta: {
title: 'Signin',
},
},

{
path: '/signup',
name: 'Signup',
component: () => import('../views/Auth/Signup.vue'),
meta: {
title: 'Signup',
},
},
],
})

export default router

router.beforeEach((to, from, next) => {
document.title = `Vue.js ${to.meta.title} | TailAdmin - Vue.js Tailwind CSS Dashboard Template`
next()
})

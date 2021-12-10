import axiosClient from '@aspida/axios'
import api from "../api/$api"
import axios from 'axios'

export const aCl = axios.create({})
export const apiClient = api(axiosClient())

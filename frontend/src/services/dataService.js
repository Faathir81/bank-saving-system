import api from './api';

export const dataService = {
  getDashboardData: async () => {
    const [custRes, accRes, depRes] = await Promise.all([
      api.get('/customers'),
      api.get('/accounts'),
      api.get('/deposito-types')
    ]);
    return {
      customers: custRes.data,
      accounts: accRes.data,
      depositoTypes: depRes.data
    };
  },

  createCustomer: async (name) => {
    return await api.post('/customers', { name });
  },

  createAccount: async (customerId, depositoTypeId) => {
    return await api.post('/accounts', { 
      customer_id: customerId, 
      deposito_type_id: depositoTypeId 
    });
  },

  deposit: async (accountId, amount, date) => {
    return await api.post('/transactions/deposit', { 
      account_id: accountId, 
      amount: parseFloat(amount), 
      date 
    });
  },

  withdraw: async (accountId, amount, date) => {
    return await api.post('/transactions/withdraw', { 
      account_id: accountId, 
      amount: parseFloat(amount), 
      date 
    });
  }
};

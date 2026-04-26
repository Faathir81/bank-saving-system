import api from './api';

export const dataService = {
  getCustomers: async () => api.get('/customers'),
  getAccounts: async () => api.get('/accounts'),
  getDepositoTypes: async () => api.get('/deposito-types'),


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
  },

  updateCustomer: async (id, name) => {
    return await api.put(`/customers/${id}`, { name });
  },

  deleteCustomer: async (id) => {
    return await api.delete(`/customers/${id}`);
  },

  updateAccount: async (id, depositoTypeId) => {
    return await api.put(`/accounts/${id}`, { deposito_type_id: depositoTypeId });
  },

  deleteAccount: async (id) => {
    return await api.delete(`/accounts/${id}`);
  },

  createDepositoType: async (name, yearly_return) => {
    return await api.post('/deposito-types', { name, yearly_return: parseFloat(yearly_return) });
  },

  updateDepositoType: async (id, name, yearly_return) => {
    return await api.put(`/deposito-types/${id}`, { name, yearly_return: parseFloat(yearly_return) });
  },

  deleteDepositoType: async (id) => {
    return await api.delete(`/deposito-types/${id}`);
  }
};

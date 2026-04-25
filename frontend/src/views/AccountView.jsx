import React, { useState } from 'react';
import { Plus, CreditCard, ArrowUpRight, ArrowDownLeft, Wallet } from 'lucide-react';
import { dataService } from '../services/dataService';

function AccountView({ accounts, customers, depositoTypes, onRefresh }) {
  const [selectedCustomer, setSelectedCustomer] = useState('');
  const [selectedType, setSelectedType] = useState('');
  const [transactionAccount, setTransactionAccount] = useState(null);
  const [amount, setAmount] = useState('');
  const [date, setDate] = useState(new Date().toISOString().split('T')[0]);
  const [calcResult, setCalcResult] = useState(null);

  const handleOpenAccount = async (e) => {
    e.preventDefault();
    if (!selectedCustomer || !selectedType) return;
    await dataService.createAccount(selectedCustomer, selectedType);
    onRefresh();
  };

  const handleTransaction = async (type) => {
    try {
      let res;
      if (type === 'deposit') {
        res = await dataService.deposit(transactionAccount.id, amount, date);
      } else {
        res = await dataService.withdraw(transactionAccount.id, amount, date);
        setCalcResult(res.data); // Store calculation result for withdrawal
      }
      if (type === 'deposit') {
        alert("Deposit Successful!");
        setTransactionAccount(null);
        setAmount('');
        onRefresh();
      }
    } catch (err) {
      alert(err.response?.data?.message || "Transaction failed");
    }
  };

  return (
    <div className="flex flex-col gap-8 pb-20">
      {/* Create Account Form */}
      <div className="card max-w-2xl">
        <h2 className="text-xl font-bold mb-4 flex items-center gap-2">
          <Plus size={20} className="text-primary"/> Open New Account
        </h2>
        <form onSubmit={handleOpenAccount} className="flex flex-col md:flex-row gap-4">
          <select 
            className="flex-1 bg-white/5 border border-white/10 rounded-lg p-3 outline-none focus:border-primary text-slate-200 min-w-[200px]"
            value={selectedCustomer}
            onChange={(e) => setSelectedCustomer(e.target.value)}
          >
            <option value="" className="bg-background">Select Customer</option>
            {customers.map(c => <option key={c.id} value={c.id} className="bg-background">{c.name}</option>)}
          </select>

          <select 
            className="flex-1 bg-white/5 border border-white/10 rounded-lg p-3 outline-none focus:border-primary text-slate-200 min-w-[200px]"
            value={selectedType}
            onChange={(e) => setSelectedType(e.target.value)}
          >
            <option value="" className="bg-background">Select Deposito Package</option>
            {depositoTypes.map(t => <option key={t.id} value={t.id} className="bg-background">{t.name} ({t.yearly_return * 100}%)</option>)}
          </select>

          <button type="submit" className="btn-primary px-8 whitespace-nowrap">Open Account</button>
        </form>
      </div>

      {/* Account List */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {accounts.map(acc => (
          <div key={acc.id} className="card group relative overflow-hidden">
            {/* Background Decor */}
            <div className="absolute -top-6 -right-6 p-4 opacity-5 group-hover:opacity-10 transition-all rotate-12">
               <Wallet size={120} />
            </div>
            
            <div className="relative z-10 flex justify-between items-start mb-8">
              <div>
                <p className="text-primary text-[10px] font-bold mb-1 uppercase tracking-[0.2em]">{acc.deposito_type.name} ACCOUNT</p>
                <h3 className="text-2xl font-bold text-white">{acc.customer.name}</h3>
                <p className="text-secondary text-[10px] font-mono mt-1 opacity-50">ID: {acc.id}</p>
              </div>
              <div className="text-right">
                <p className="text-secondary text-xs font-medium mb-1">Current Balance</p>
                <p className="text-3xl font-black text-white tracking-tight">
                  <span className="text-primary mr-1 text-lg">Rp</span>
                  {acc.balance.toLocaleString()}
                </p>
              </div>
            </div>

            <div className="relative z-10 flex gap-3">
              <button 
                onClick={() => {setTransactionAccount(acc); setCalcResult(null);}}
                className="flex-1 bg-white/5 hover:bg-white/10 p-3 rounded-lg flex items-center justify-center gap-2 transition-all"
              >
                <ArrowUpRight size={18} className="text-emerald-400"/> Deposit
              </button>
              <button 
                onClick={() => {setTransactionAccount(acc); setCalcResult(null);}}
                className="flex-1 bg-white/5 hover:bg-white/10 p-3 rounded-lg flex items-center justify-center gap-2 transition-all"
              >
                <ArrowDownLeft size={18} className="text-accent"/> Withdraw
              </button>
            </div>
          </div>
        ))}
      </div>

      {/* Transaction Modal (Simplified for demo) */}
      {transactionAccount && (
        <div className="fixed inset-0 bg-background/80 backdrop-blur-sm flex items-center justify-center z-50 p-6">
          <div className="card w-full max-w-md shadow-2xl border-primary/20">
            <h2 className="text-2xl font-bold mb-2">Process Transaction</h2>
            <p className="text-secondary mb-6">Account: {transactionAccount.customer.name}</p>
            
            <div className="flex flex-col gap-4">
              <div className="flex flex-col gap-1">
                <label className="text-xs text-secondary ml-1">Amount (Rp)</label>
                <input 
                  type="number" 
                  placeholder="0" 
                  className="bg-white/5 border border-white/10 rounded-lg p-4 text-xl font-bold outline-none focus:border-primary"
                  value={amount}
                  onChange={(e) => setAmount(e.target.value)}
                />
              </div>

              <div className="flex flex-col gap-1">
                <label className="text-xs text-secondary ml-1">Date</label>
                <input 
                  type="date" 
                  className="bg-white/5 border border-white/10 rounded-lg p-3 outline-none focus:border-primary"
                  value={date}
                  onChange={(e) => setDate(e.target.value)}
                />
              </div>

              {calcResult && (
                <div className="bg-emerald-500/10 border border-emerald-500/20 rounded-xl p-4 mt-2">
                   <p className="text-xs text-emerald-400 font-bold uppercase mb-2">Withdrawal Result</p>
                   <div className="flex justify-between text-sm mb-1">
                      <span>Interest Earned ({calcResult.months_stayed} mo):</span>
                      <span className="font-bold">Rp {calcResult.interest_earned.toLocaleString()}</span>
                   </div>
                   <div className="flex justify-between text-lg font-bold border-t border-emerald-500/20 pt-2 mt-2">
                      <span>Total Received:</span>
                      <span className="text-emerald-400">Rp {calcResult.total_received.toLocaleString()}</span>
                   </div>
                </div>
              )}

              <div className="flex gap-3 mt-4">
                <button 
                  onClick={() => setTransactionAccount(null)}
                  className="flex-1 bg-white/5 p-4 rounded-xl font-bold"
                >
                  Cancel
                </button>
                {!calcResult ? (
                  <>
                    <button onClick={() => handleTransaction('deposit')} className="flex-1 bg-emerald-600 p-4 rounded-xl font-bold">Deposit</button>
                    <button onClick={() => handleTransaction('withdraw')} className="flex-1 bg-accent p-4 rounded-xl font-bold">Withdraw</button>
                  </>
                ) : (
                  <button onClick={() => {setTransactionAccount(null); onRefresh();}} className="flex-1 bg-primary p-4 rounded-xl font-bold">Done</button>
                )}
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default AccountView;

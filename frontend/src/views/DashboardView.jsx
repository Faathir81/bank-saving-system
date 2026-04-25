import React from 'react';
import { Users, CreditCard, TrendingUp } from 'lucide-react';

function DashboardView({ customers, accounts }) {
  const totalBalance = accounts.reduce((acc, curr) => acc + curr.balance, 0);

  return (
    <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
      <StatCard 
        label="Total Balance" 
        value={`Rp ${totalBalance.toLocaleString()}`} 
        icon={<TrendingUp className="text-emerald-400"/>} 
      />
      <StatCard 
        label="Active Customers" 
        value={customers.length} 
        icon={<Users className="text-blue-400"/>} 
      />
      <StatCard 
        label="Active Accounts" 
        value={accounts.length} 
        icon={<CreditCard className="text-purple-400"/>} 
      />
    </div>
  );
}

function StatCard({ label, value, icon }) {
  return (
    <div className="card flex items-center justify-between">
      <div>
        <p className="text-secondary text-sm mb-1">{label}</p>
        <h3 className="text-2xl font-bold">{value}</h3>
      </div>
      <div className="p-3 bg-white/5 rounded-lg">
        {icon}
      </div>
    </div>
  );
}

export default DashboardView;

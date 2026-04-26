import React from 'react';
import { LayoutDashboard, Users, CreditCard, TrendingUp, Package } from 'lucide-react';

function Sidebar({ activeTab, setActiveTab }) {
  return (
    <aside className="w-64 glass border-r border-white/10 p-6 flex flex-col gap-8">
      <div className="flex items-center gap-3 px-2">
        <div className="w-10 h-10 bg-primary rounded-xl flex items-center justify-center shadow-lg shadow-primary/20">
          <TrendingUp size={24} className="text-white" />
        </div>
        <span className="text-xl font-bold tracking-tight text-white">BankSave</span>
      </div>

      <nav className="flex flex-col gap-2">
        <NavItem 
          active={activeTab === 'dashboard'} 
          onClick={() => setActiveTab('dashboard')} 
          icon={<LayoutDashboard size={20}/>} 
          label="Dashboard" 
        />
        <NavItem 
          active={activeTab === 'customers'} 
          onClick={() => setActiveTab('customers')} 
          icon={<Users size={20}/>} 
          label="Customers" 
        />
        <NavItem 
          active={activeTab === 'accounts'} 
          onClick={() => setActiveTab('accounts')} 
          icon={<CreditCard size={20}/>} 
          label="Accounts" 
        />
        <NavItem 
          active={activeTab === 'packages'} 
          onClick={() => setActiveTab('packages')} 
          icon={<Package size={20}/>} 
          label="Packages" 
        />
      </nav>
    </aside>
  );
}

function NavItem({ icon, label, active, onClick }) {
  return (
    <button 
      onClick={onClick}
      className={`flex items-center gap-3 px-4 py-3 rounded-xl transition-all ${
        active ? 'bg-primary text-white shadow-lg shadow-primary/20' : 'hover:bg-white/5 text-secondary hover:text-white'
      }`}
    >
      {icon}
      <span className="font-medium">{label}</span>
    </button>
  );
}

export default Sidebar;

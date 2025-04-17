'use client';

import {PrivyProvider} from '@privy-io/react-auth';
import {b3Sepolia, baseSepolia, optimismSepolia} from "viem/chains";

export default function PrivyAuthProvider({children}: {children: React.ReactNode}) {
    return (
        <PrivyProvider
            appId={process.env.NEXT_PUBLIC_PRIVY_APP_ID as string}
            clientId={process.env.NEXT_PUBLIC_PRIVY_CLIENT_ID as string}
            config={{
                // Customize Privy's appearance in your app
                appearance: {
                    theme: 'dark',
                    accentColor: '#676FFF',
                    landingHeader: "DeFree",
                    loginMessage: "Connect to the world of decentralized work!",
                    showWalletLoginFirst: true
                },
                embeddedWallets: {
                    createOnLogin: 'users-without-wallets'
                },
                defaultChain: optimismSepolia,
                supportedChains: [optimismSepolia, baseSepolia, b3Sepolia]
            }}
        >
            {children}
        </PrivyProvider>
    );
}
import { useEffect } from 'react';

const TelegramLogin = () => {
    useEffect(() => {
        window.TelegramLoginWidget = {
            dataOnAuth: (user: any) => {
                fetch('/api/auth/telegram', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(user),
                })
                    .then((res) => res.json())
                    .then((data) => {
                        if (data.token) {
                            localStorage.setItem('token', data.token);
                        } else {
                            console.error('Authentication failed');
                        }
                    });
            },
        };
    }, []);

    return (
        <div>
            <script
                async
                src={`https://telegram.org/js/telegram-widget.js?15`}
                data-telegram-login="your_bot_username"
                data-size="large"
                data-radius="10"
                data-auth-url="https://your-backend.com/api/auth/telegram"
                data-request-access="write"
            />
        </div>
    );
};

export default TelegramLogin;

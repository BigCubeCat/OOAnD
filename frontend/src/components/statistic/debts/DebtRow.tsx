import {Box, Avatar,Typography} from "@mui/joy";


interface IDebtRowProps {
    icon: string;
    name: string;
    amount: number;
}

export default function DebtRow(props: IDebtRowProps) {
    const isGreen = props.amount > 0;
    const color = (isGreen) ? 'success.100' : 'danger.100';
    const prefix = isGreen ? '+' : '';
    return (
        <Box sx={{
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center',
            marginTop: 1
        }}>
            <Box sx={{display: 'flex', alignItems: 'center'}}>
            <Avatar size='lg' src={props.icon} />
            <Typography sx={{marginLeft: 2}}>{props.name}</Typography>
            </Box>
            <Box sx={{color: color, fontWeight: 1000}}>{prefix}{props.amount}</Box>
        </Box>
    )
}
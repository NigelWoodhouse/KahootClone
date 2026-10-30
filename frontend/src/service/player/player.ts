import { writable, type Writable } from "svelte/store";
import { NetService, PacketTypes, type Packet, type ConnectPacket, GameState, type ChangeGameStatePacket, type QuestionAnswerPacket, type PlayerRevealPacket, type LeaderboardPacket, type LeaderboardEntry } from "../net";

export const gameState: Writable<GameState> = writable(GameState.Lobby);
export const state = gameState;
export const points: Writable<number> = writable(0);
export const leaderboard: Writable<LeaderboardEntry[]> = writable([]);
export const intermissionImage: Writable<string | null> = writable(null);
export const intermissionQuestionName: Writable<string | null> = writable(null);
export const gameCancelled: Writable<boolean> = writable(false);

export class PlayerGame {
    private net: NetService;

    constructor(){
        this.net = new NetService();
        this.net.connect();
        this.net.onPacket(p => this.onPacket(p));
    }

    join(code: string, name: string){
        gameCancelled.set(false);
        let packet: ConnectPacket = {
            id: PacketTypes.Connect,
            code: code,
            name: name,
        }

        this.net.sendPacket(packet);
    }

    answer(question: number){
        let packet: QuestionAnswerPacket = {
            id: PacketTypes.Answer,
            question: question
        };

        this.net.sendPacket(packet);
    }

    onPacket(packet: Packet){
        switch(packet.id){
            case PacketTypes.CancelGame: {
                gameCancelled.set(true);
                gameState.set(GameState.Lobby);
                points.set(0);
                leaderboard.set([]);
                intermissionImage.set(null);
                intermissionQuestionName.set(null);
                break;
            }

            case PacketTypes.ChangeGameState:{
                let data = packet as ChangeGameStatePacket;
                gameState.set(data.state);
                break;
            }
            case PacketTypes.PlayerReveal:{
                let data = packet as PlayerRevealPacket;
                points.set(data.points);
                break;
            }
            case PacketTypes.Leaderboard:{
                let data = packet as LeaderboardPacket;
                leaderboard.set(data.points);
                if (data.image !== undefined) {
                    intermissionImage.set(data.image || null);
                }
                if (data.questionName !== undefined) {
                    intermissionQuestionName.set(data.questionName || null);
                }
                break;
            }
        }
    }
}
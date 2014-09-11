 {- 
  - http://sequence.complete.org/node/258
  - https://github.com/haskell/network/blob/master/examples/EchoServer.hs
  - https://hackage.haskell.org/package/network
  -
  -}

module Main where

import Control.Monad (unless)
import Network.Socket hiding (recv)
import qualified Data.ByteString as S
import qualified Data.ByteString.Char8 as BC
import Network.Socket.ByteString (recv, sendAll)

main :: IO()
main = withSocketsDo $
    do addrinfos <- getAddrInfo
                    (Just (defaultHints {addrFlags = [AI_PASSIVE]}))
                    Nothing (Just "3333")
       let serverAddr = head addrinfos
       sock <- socket (addrFamily serverAddr) Stream defaultProtocol
       bindSocket sock (addrAddress serverAddr)
       listen sock 100
       talk   sock
       sClose sock

    where
        talk :: Socket -> IO ()
        talk sock = 
            do (conn, _) <- accept sock
               msg <- recv conn 1024
               unless (BC.null msg) $ sendAll conn msg >> BC.putStrLn msg
               talk sock 
